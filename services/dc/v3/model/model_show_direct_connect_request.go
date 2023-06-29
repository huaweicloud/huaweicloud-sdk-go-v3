package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDirectConnectRequest Request Object
type ShowDirectConnectRequest struct {

	// 物理专线连接ID。
	DirectConnectId string `json:"direct_connect_id"`

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条资源记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

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
