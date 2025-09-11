package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIpdProjectIssueResponse Response Object
type CreateIpdProjectIssueResponse struct {

	// 返回信息
	Message *string `json:"message,omitempty"`

	// 返回对象列表
	Result *[]interface{} `json:"result,omitempty"`

	// 返回状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateIpdProjectIssueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpdProjectIssueResponse struct{}"
	}

	return strings.Join([]string{"CreateIpdProjectIssueResponse", string(data)}, " ")
}
