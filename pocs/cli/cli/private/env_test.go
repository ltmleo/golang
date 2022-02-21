package private

import (
	"testing"
)

func TestGitEnv_Load(t *testing.T) {
	fileName := "mock/test.yaml"
	env := Env{
		file: fileName,
	}
	env.Load()
}