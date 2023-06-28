package lib

import (
	"github.com/golang/protobuf/proto"
	"github.com/unity-sds/unity-cs-manager/internal/pkg/eks"
	"github.com/unity-sds/unity-cs-manager/marketplace"
)

func GenerateEKSTemplate(arr []byte) (string, error) {
	model := &marketplace.Install_Extensions_Eks{}
	err := proto.Unmarshal(arr, model)
	if err != nil {
		return "", err
	}
	return eks.ProtoGenerate(*model)
}
