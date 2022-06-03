package k8s

import (
	"context"
	"os"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type DeploymentList *appsv1.DeploymentList
type Config *rest.Config

func setConfig(configMode string) (Config, error) {
	if configMode == "in" {
		config, err := rest.InClusterConfig()
		return config, err
	}
	kubeconfig := os.Getenv("KUBECONFIG")
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	return config, err
}

func List(namespace string) DeploymentList {
	configMode := os.Getenv("K8S_CONFIG_MODE")
	config, err := setConfig(configMode)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return deployments
}