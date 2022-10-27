package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Archive struct {

	// 镜像命名空间。
	ArtifactNamespace *string `json:"artifact_namespace,omitempty"`
}

func (o Archive) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Archive struct{}"
	}

	return strings.Join([]string{"Archive", string(data)}, " ")
}
