package esteira

import (
	"encoding/json"

	"github.com/ltmleo/golang/kubernetes/copy-ns/internal/git"
	"github.com/ltmleo/golang/kubernetes/copy-ns/internal/k8s"
)

type Release struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Esteira map[string]Release

func (e Esteira) Load(deployments k8s.DeploymentList) {
	r := Release{}
	for _, d := range deployments.Items {
		r = Release{Name: d.Name, Image: d.Spec.Template.Spec.Containers[0].Image}
		e[d.Name] = r
	}
}

func (e Esteira) ToJSON() []byte {
	j, _ := json.MarshalIndent(e, "", "\t")
	return j
}

func (e *Esteira) GetFromGit(projectID string, fileName string) {
	fileSha := git.GetFileSha(projectID, fileName, "", "esteira")
	esteira := git.GetFile(projectID, fileSha)
	json.Unmarshal([]byte(esteira), &e)
}
