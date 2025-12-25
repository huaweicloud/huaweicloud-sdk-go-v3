package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulnerabilitiesResponse Response Object
type ListVulnerabilitiesResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	// 漏洞总数
	Total *int32 `json:"total,omitempty"`

	// 分页大小
	Size *int32 `json:"size,omitempty"`

	// 偏移量
	Page *int32 `json:"page,omitempty"`

	// 是否成功
	Success *bool `json:"success,omitempty"`

	// 漏洞列表
	Data           *[]VulnerabilityDetail `json:"data,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListVulnerabilitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulnerabilitiesResponse struct{}"
	}

	return strings.Join([]string{"ListVulnerabilitiesResponse", string(data)}, " ")
}
