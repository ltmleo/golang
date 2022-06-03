package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ltmleo/golang/kubernetes/copy-ns/internal/esteira"
	"github.com/ltmleo/golang/kubernetes/copy-ns/internal/git"
	"github.com/ltmleo/golang/kubernetes/copy-ns/internal/helm"
	"github.com/ltmleo/golang/kubernetes/copy-ns/internal/k8s"
)

func main() {
	env := os.Getenv("ENVIRONMENT")
	if "" == env {
		env = "development"
	}
	if "production" != env {
		err := godotenv.Load(".env." + env)
		if err != nil {
			panic(err)
		}
	}
	//getEsteira("dev")
	applyEsteira("dev.json", "dev")

}

func getEsteira(namespace string) {
	e := make(esteira.Esteira)
	deploys := k8s.List(namespace)
	e.Load(deploys)
	fmt.Println(string(e.ToJSON()))
	actions := []git.Action{{Action: "update", FilePath: "/esteira/" + namespace + ".json", Content: string(e.ToJSON())}}
	if os.Getenv("COMMIT_CHANGES") == "true" {
		result := git.Commit(actions)
		if strings.Contains(result, "A file with this name doesn't exist") {
			actions[0].Action = "create"
			result = git.Commit(actions)
		}
		fmt.Println(result)
	}
}

func applyEsteira(fileName string, destiny string) {
	blackList := os.Getenv("RELEASES_BLACK_LIST")
	esteira := esteira.Esteira{}
	esteira.GetFromGit(os.Getenv("GITLAB_GROUP_ID"), fileName)
	for name, info := range esteira {
		if !strings.Contains(blackList, name) {
			values, _ := git.GetEnvFile(name, destiny)
			versionArray := strings.Split(info.Image, ":")
			version := versionArray[len(versionArray)-1]
			fmt.Println(values)
			fmt.Println(helm.HelmInstall([]byte(values), name, version, destiny))
		}
	}

}
