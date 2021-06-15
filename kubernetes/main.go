/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"
    "os"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

var clientset *kubernetes.Clientset

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(os.Getenv("KUBECONFIG")), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	GetPods("istio-system")
	GetPods("")
}

func GetPods(namespace string) {
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the namespace %s\n", len(pods.Items), namespace)

	for _,pod := range pods.Items {
		fmt.Printf("Namespace: %s, Name: %s\n", pod.Namespace, pod.Name)
	}
	CreateConfiMap()
}

func CreateConfiMap() {
	configMapClient := clientset.CoreV1().ConfigMaps(apiv1.NamespaceDefault)
	configMap := &apiv1.ConfigMap{
					ObjectMeta: metav1.ObjectMeta{
						Name: "config-example",
					},
					Data: map[string]string{
						"app": "demo",
					},
				}
	res, err := configMapClient.Create(context.TODO(), configMap, metav1.CreateOptions{})
	if err != nil {
		fmt.Println("Delete configMap config-example")
		configMapClient.Delete(context.TODO(), "config-example", metav1.DeleteOptions{})
	}
	fmt.Printf("Created ConfigMap %q.\n", res.GetObjectMeta().GetName())
	
}