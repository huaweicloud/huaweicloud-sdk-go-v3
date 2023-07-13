package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNodeActivationRecordsRequest Request Object
type ShowNodeActivationRecordsRequest struct {

	// 设备ID，从专业版HiLens控制台设备管理[查询设备列表](ListNodeUsingGeT.xml)获取
	NodeId string `json:"node_id"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，取值范围1~100，默认为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowNodeActivationRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeActivationRecordsRequest struct{}"
	}

	return strings.Join([]string{"ShowNodeActivationRecordsRequest", string(data)}, " ")
}
