package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrashArtifactModelForDelete struct {

	// 仓库id
	Id string `json:"id"`

	// 仓库类型
	Fomat string `json:"fomat"`

	// URI
	Uri string `json:"uri"`

	// 状态
	Status string `json:"status"`

	// 路径白名单
	IncludePattern *string `json:"include_pattern,omitempty"`
}

func (o TrashArtifactModelForDelete) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrashArtifactModelForDelete struct{}"
	}

	return strings.Join([]string{"TrashArtifactModelForDelete", string(data)}, " ")
}
