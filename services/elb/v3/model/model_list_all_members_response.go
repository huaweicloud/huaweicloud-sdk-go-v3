package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAllMembersResponse struct {
	// 请求ID。  注：自动生成 。
	RequestId *string   `json:"request_id,omitempty"`
	PageInfo  *PageInfo `json:"page_info,omitempty"`
	// 后端服务器对象列表。
	Members        *[]Member `json:"members,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAllMembersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAllMembersResponse struct{}"
	}

	return strings.Join([]string{"ListAllMembersResponse", string(data)}, " ")
}
