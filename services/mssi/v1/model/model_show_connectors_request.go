package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectorsRequest Request Object
type ShowConnectorsRequest struct {

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset int32 `json:"offset"`

	// 每页显示的条目数量
	Limit int32 `json:"limit"`

	// 连接器分类:onboard：查询所有的上架的、custom：自己创建的连接器
	Scope *string `json:"scope,omitempty"`

	// 连接器种类（连接器市场的tab分类）
	Category *string `json:"category,omitempty"`

	// CustomConnectors的名字
	Name *string `json:"name,omitempty"`

	// operation条件过滤
	OperationType *string `json:"operation_type,omitempty"`
}

func (o ShowConnectorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectorsRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectorsRequest", string(data)}, " ")
}
