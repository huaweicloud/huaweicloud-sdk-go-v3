package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAllMembersResponse struct {

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty" xml:"page_info"`

	// 后端服务器对象列表。
	Members        *[]Member `json:"members,omitempty" xml:"members"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAllMembersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllMembersResponse struct{}"
	}

	return strings.Join([]string{"ListAllMembersResponse", string(data)}, " ")
}
