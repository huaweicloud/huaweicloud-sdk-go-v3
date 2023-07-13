package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEffectiveRoutesResponse Response Object
type ListEffectiveRoutesResponse struct {

	// 路由列表
	Routes *[]EffectiveRoute `json:"routes,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListEffectiveRoutesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEffectiveRoutesResponse struct{}"
	}

	return strings.Join([]string{"ListEffectiveRoutesResponse", string(data)}, " ")
}
