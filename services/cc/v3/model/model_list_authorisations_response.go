package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuthorisationsResponse Response Object
type ListAuthorisationsResponse struct {

	// 授权的详细信息。
	Authorisations *[]Authorisation `json:"authorisations,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAuthorisationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorisationsResponse struct{}"
	}

	return strings.Join([]string{"ListAuthorisationsResponse", string(data)}, " ")
}
