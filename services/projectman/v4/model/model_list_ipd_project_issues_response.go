package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpdProjectIssuesResponse Response Object
type ListIpdProjectIssuesResponse struct {

	// 返回信息
	Message *string `json:"message,omitempty"`

	// 返回对象
	Result *interface{} `json:"result,omitempty"`

	// 返回状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListIpdProjectIssuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpdProjectIssuesResponse struct{}"
	}

	return strings.Join([]string{"ListIpdProjectIssuesResponse", string(data)}, " ")
}
