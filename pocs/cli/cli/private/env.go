package private

import (
	"helm.sh/helm/v3/pkg/chart/loader"
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chartutil"

)


func Execute() {
	c, _ := loader.Load("/home/vivo/Documents/VIVO/iac-lojaonline-envs/chart/")
	vals, _ := chartutil.ReadValuesFile("/home/vivo/Documents/VIVO/iac-lojaonline-envs/agreement.yaml") 
	cfg := new(action.Configuration)
	client := action.NewInstall(cfg)
	client.DryRun = true
	client.ReleaseName = "agreement-env"
	client.Replace = true // Skip the name check
	client.ClientOnly = true
	client.IncludeCRDs = false
	client.Namespace = "dev"
	rel, _ := client.Run(c, vals)
	fmt.Println(rel.Manifest)

}
