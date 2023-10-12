package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompatibleInstanceVolumeResp struct {

	// 类型
	Type *string `json:"type,omitempty"`

	// 连接
	Used *float32 `json:"used,omitempty"`

	// 大小
	Size *int32 `json:"size,omitempty"`
}

func (o CompatibleInstanceVolumeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompatibleInstanceVolumeResp struct{}"
	}

	return strings.Join([]string{"CompatibleInstanceVolumeResp", string(data)}, " ")
}
