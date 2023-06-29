package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssociationsResponse Response Object
type ListAssociationsResponse struct {

	// 路由表关联列表
	Associations *[]Association `json:"associations,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAssociationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociationsResponse struct{}"
	}

	return strings.Join([]string{"ListAssociationsResponse", string(data)}, " ")
}
