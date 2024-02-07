package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNewNodeAutoAddSwitchRequestBody 开启或关闭新增节点自动加入该Proxy请求体。
type UpdateNewNodeAutoAddSwitchRequestBody struct {

	// 是否开启新增节点自动加入该Proxy。  取值范围： - ON：开启。 - OFF：关闭。
	SwitchStatus string `json:"switch_status"`

	// 新增节点的读权重：  - 新增节点自动加入为ON，取值为0~1000。 - 新增节点自动加入为OFF，则可不输入读权重。
	Weight *int32 `json:"weight,omitempty"`
}

func (o UpdateNewNodeAutoAddSwitchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNewNodeAutoAddSwitchRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNewNodeAutoAddSwitchRequestBody", string(data)}, " ")
}
