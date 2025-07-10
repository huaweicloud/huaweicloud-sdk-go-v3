package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIssueDetailResponse Response Object
type ShowIssueDetailResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 信息
	Message *string `json:"message,omitempty"`

	// 查询结果
	Result         *[]map[string]interface{} `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowIssueDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIssueDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowIssueDetailResponse", string(data)}, " ")
}
