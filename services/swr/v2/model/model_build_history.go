package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BuildHistory struct {

	// 构建时间
	CreatedAt string `json:"created_at"`

	// 构建命令
	CreatedBy string `json:"created_by"`

	// 是否空层
	EmptyLayer bool `json:"empty_layer"`

	// 镜像层格式
	MediaType string `json:"mediaType"`

	// 镜像层大小
	Size int32 `json:"size"`

	// 镜像层sha256信息
	Digest string `json:"digest"`
}

func (o BuildHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildHistory struct{}"
	}

	return strings.Join([]string{"BuildHistory", string(data)}, " ")
}
