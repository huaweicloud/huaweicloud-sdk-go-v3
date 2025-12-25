package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataclassFieldsResponse Response Object
type ListDataclassFieldsResponse struct {

	// 字段列表详情
	FieldDetails *[]FieldResponseBody `json:"field_details,omitempty"`

	// 数据总量
	Total float32 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDataclassFieldsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataclassFieldsResponse struct{}"
	}

	return strings.Join([]string{"ListDataclassFieldsResponse", string(data)}, " ")
}
