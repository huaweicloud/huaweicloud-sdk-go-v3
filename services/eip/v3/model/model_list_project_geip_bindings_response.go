package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectGeipBindingsResponse Response Object
type ListProjectGeipBindingsResponse struct {

	// geip绑定关系对象
	GeipBindings *[]GeipBindingsInternalResp `json:"geip_bindings,omitempty"`

	// 本次请求编号
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectGeipBindingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectGeipBindingsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectGeipBindingsResponse", string(data)}, " ")
}
