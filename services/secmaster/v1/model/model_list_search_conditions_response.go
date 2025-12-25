package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSearchConditionsResponse Response Object
type ListSearchConditionsResponse struct {

	// 检索条件个数
	Count *int64 `json:"count,omitempty"`

	// 检索条件集合
	Records        *[]ListSearchConditionsDetail `json:"records,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListSearchConditionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchConditionsResponse struct{}"
	}

	return strings.Join([]string{"ListSearchConditionsResponse", string(data)}, " ")
}
