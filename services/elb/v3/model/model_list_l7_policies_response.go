package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListL7PoliciesResponse Response Object
type ListL7PoliciesResponse struct {

	// 参数解释：请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 参数解释：转发策略对象列表。
	L7policies     *[]L7Policy `json:"l7policies,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListL7PoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListL7PoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListL7PoliciesResponse", string(data)}, " ")
}
