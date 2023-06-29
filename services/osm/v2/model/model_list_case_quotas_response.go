package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCaseQuotasResponse Response Object
type ListCaseQuotasResponse struct {

	// 总配额
	Total *int32 `json:"total,omitempty"`

	// 未使用
	UnUsed         *int32 `json:"un_used,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCaseQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaseQuotasResponse struct{}"
	}

	return strings.Join([]string{"ListCaseQuotasResponse", string(data)}, " ")
}
