package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRemediationExceptionsResponse Response Object
type ListRemediationExceptionsResponse struct {

	// 合规规则修正例外的详情。
	Value *[]RemediationException `json:"value,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRemediationExceptionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRemediationExceptionsResponse struct{}"
	}

	return strings.Join([]string{"ListRemediationExceptionsResponse", string(data)}, " ")
}
