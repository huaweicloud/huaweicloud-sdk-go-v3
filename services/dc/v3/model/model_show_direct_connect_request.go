package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDirectConnectRequest Request Object
type ShowDirectConnectRequest struct {

	// 物理专线连接ID。
	DirectConnectId string `json:"direct_connect_id"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`
}

func (o ShowDirectConnectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDirectConnectRequest struct{}"
	}

	return strings.Join([]string{"ShowDirectConnectRequest", string(data)}, " ")
}
