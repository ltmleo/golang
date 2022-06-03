package helm

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/cli"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/rest"
)

var kubeConfig *genericclioptions.ConfigFlags

func HelmInstall(values []byte, releaseName string, version string, namespace string) string{
	cfg := new(action.Configuration)
	client := action.NewUpgrade(cfg)
	//TODO: criptografar
	charURL := "http://"+os.Getenv("NEXUS_USER")+":"+os.Getenv("NEXUS_PASSWORD")+"@"+os.Getenv("NEXUS_HELM_CHART_URL")
	chart_path, err := client.LocateChart(charURL, cli.New())
	if err != nil {
        panic(err)
	}
	c, err := loader.Load(chart_path)
	if err != nil {
			panic(err)
	}
	values = []byte(string(values) + "\nimage:\n  tag: "+version)
	vals, err := chartutil.ReadValues(values)
	if err != nil {
			panic(err)
	}
	client.DryRun, _ = strconv.ParseBool(os.Getenv("HELM_DRY_RUN"))

	kubeConfig := getKubeConfig(namespace)
	if err := cfg.Init(kubeConfig, namespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	}); err != nil {
		panic(err)
	}
	client.Force, _ = strconv.ParseBool(os.Getenv("HELM_UPGRADE_FORCE"))
	client.Install = true
	client.Namespace = namespace
	rel, err := client.Run(releaseName, c, vals)
	if err != nil {
		if strings.Contains(err.Error(), "has no deployed releases") {
			clientInstall := action.NewInstall(cfg)
			clientInstall.DryRun = client.DryRun
			clientInstall.ReleaseName = releaseName
			clientInstall.Namespace = namespace
			rel, err = clientInstall.Run(c, vals)
			if err != nil {
				panic(err)
			}
		}
	}
	if client.DryRun{
		return rel.Manifest
	}
	return rel.Info.Description

}

func getKubeConfig(namespace string)  *genericclioptions.ConfigFlags {
	if os.Getenv("K8S_CONFIG_MODE") == "in"{
		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err)
		}
		kubeConfig = genericclioptions.NewConfigFlags(false)
		kubeConfig.APIServer = &config.Host
		kubeConfig.BearerToken = &config.BearerToken
		kubeConfig.CAFile = &config.CAFile
		kubeConfig.Namespace = &namespace
	} else {
		kubeConfig = kube.GetConfig(os.Getenv("KUBECONFIG"), "", namespace)
	}
	return kubeConfig
}