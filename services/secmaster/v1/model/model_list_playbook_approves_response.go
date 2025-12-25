package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookApprovesResponse Response Object
type ListPlaybookApprovesResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应消息
	Message *string `json:"message,omitempty"`

	// 剧本审核详情
	Data *[]ApproveOpinionDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlaybookApprovesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookApprovesResponse struct{}"
	}

	return strings.Join([]string{"ListPlaybookApprovesResponse", string(data)}, " ")
}
