package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNewNodeAutoAddSwitchResponse Response Object
type UpdateNewNodeAutoAddSwitchResponse struct {

	// 开启或关闭新增节点自动加入该Proxy的操作结果。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateNewNodeAutoAddSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNewNodeAutoAddSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateNewNodeAutoAddSwitchResponse", string(data)}, " ")
}
