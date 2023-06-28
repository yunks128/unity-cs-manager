package lib

import (
	"github.com/unity-sds/unity-cs-manager/internal/marketplace"
	"github.com/unity-sds/unity-cs-manager/internal/pkg/eks"
)

func GenerateEKSTemplate(model marketplace.Install_Extensions_Eks) (string, error) {
	return eks.ProtoGenerate(model)
}
