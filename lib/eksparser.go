package lib

import (
	"github.com/unity-sds/unity-cs-manager/internal/pkg/eks"
	"github.com/unity-sds/unity-cs-manager/marketplace"
)

func GenerateEKSTemplate(model *marketplace.Install_Extensions_Eks) (string, error) {
	return eks.ProtoGenerate(*model)
}
