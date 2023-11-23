package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataclassResponse Response Object
type ListDataclassResponse struct {

	// 数据类详情
	DataclassDetails *[]DataClassResponseBody `json:"dataclass_details,omitempty"`

	// 数据总量
	Total float32 `json:"total,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDataclassResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataclassResponse struct{}"
	}

	return strings.Join([]string{"ListDataclassResponse", string(data)}, " ")
}
