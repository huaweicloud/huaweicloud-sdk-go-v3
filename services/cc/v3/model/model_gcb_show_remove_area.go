package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GcbShowRemoveArea struct {

	// 功能说明：远端接入点的中英文名。通过HEADER里面的x-language控制，默认英文，zh-cn返回中文。
	RemoteArea *string `json:"remote_area,omitempty"`
}

func (o GcbShowRemoveArea) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcbShowRemoveArea struct{}"
	}

	return strings.Join([]string{"GcbShowRemoveArea", string(data)}, " ")
}
