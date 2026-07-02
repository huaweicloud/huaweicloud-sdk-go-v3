package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopVpcepResponse Response Object
type StopVpcepResponse struct {

	// 操作行为。固定为：deleteVpcepservice，表示已关闭终端节点。
	Action         *string `json:"action,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopVpcepResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopVpcepResponse struct{}"
	}

	return strings.Join([]string{"StopVpcepResponse", string(data)}, " ")
}
