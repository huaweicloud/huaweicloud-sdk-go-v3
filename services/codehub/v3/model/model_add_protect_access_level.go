package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddProtectAccessLevel struct {

	// 提交权限
	PushAccessLevel int32 `json:"push_access_level"`

	// 合并权限
	MergeAccessLevel int32 `json:"merge_access_level"`
}

func (o AddProtectAccessLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectAccessLevel struct{}"
	}

	return strings.Join([]string{"AddProtectAccessLevel", string(data)}, " ")
}
