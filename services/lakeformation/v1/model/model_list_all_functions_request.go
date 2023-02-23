package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAllFunctionsRequest struct {

	// 实例Id
	InstanceId string `json:"instance_id"`

	// catalog名字
	CatalogName string `json:"catalog_name"`

	// 查询返回条数
	Limit *int32 `json:"limit,omitempty"`

	// 查询的起始记录ID
	Marker *string `json:"marker,omitempty"`

	// 是否查询上一页
	ReversePage *bool `json:"reverse_page,omitempty"`
}

func (o ListAllFunctionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllFunctionsRequest struct{}"
	}

	return strings.Join([]string{"ListAllFunctionsRequest", string(data)}, " ")
}
