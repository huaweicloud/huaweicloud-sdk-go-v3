package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIssueTreeResponse Response Object
type ListIssueTreeResponse struct {

	// 对外时：success|error; 对内时：ok|failed
	Status *string `json:"status,omitempty"`

	Result *ResultValueListWorkItemVo `json:"result,omitempty"`

	Error *ApiError `json:"error,omitempty"`

	// 由接口调用方传入，建议使用UUID保证请求的唯一性。
	RequestId *string `json:"request_id,omitempty"`

	// 对内接口才有此属性
	ServerAddress  *string `json:"server_address,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListIssueTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIssueTreeResponse struct{}"
	}

	return strings.Join([]string{"ListIssueTreeResponse", string(data)}, " ")
}
