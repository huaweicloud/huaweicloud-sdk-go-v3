package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Archive 构建产物纳管。
type Archive struct {

	// 产物纳管SWR组织。
	ArtifactNamespace *string `json:"artifact_namespace,omitempty"`
}

func (o Archive) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Archive struct{}"
	}

	return strings.Join([]string{"Archive", string(data)}, " ")
}
