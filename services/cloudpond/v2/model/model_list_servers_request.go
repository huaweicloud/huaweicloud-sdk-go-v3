package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServersRequest Request Object
type ListServersRequest struct {

	// 边缘小站ID
	EdgeSiteId *string `json:"edge_site_id,omitempty"`

	// 每页的数量
	Limit *int32 `json:"limit,omitempty"`

	// 分页标识
	Marker *string `json:"marker,omitempty"`

	// 根据服务器状态查询，支持多值查询
	Status *[]ServerStatus `json:"status,omitempty"`

	// 根据ID过滤，支持多值查询，查询条件格式：id=xxx&id=xxx。
	Id *[]string `json:"id,omitempty"`
}

func (o ListServersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersRequest struct{}"
	}

	return strings.Join([]string{"ListServersRequest", string(data)}, " ")
}
