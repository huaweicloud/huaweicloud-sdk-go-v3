package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuthorisationsResponse Response Object
type ListAuthorisationsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 授权实例列表。
	Authorisations []Authorisation `json:"authorisations"`
	HttpStatusCode int             `json:"-"`
}

func (o ListAuthorisationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorisationsResponse struct{}"
	}

	return strings.Join([]string{"ListAuthorisationsResponse", string(data)}, " ")
}
