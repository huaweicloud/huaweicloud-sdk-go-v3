package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceResponse Response Object
type ListInstanceResponse struct {

	// 独享引擎实例数量
	Total *int32 `json:"total,omitempty"`

	// 是否曾经购买过独享引擎
	Purchased *bool `json:"purchased,omitempty"`

	// 详细的独享引擎信息列表
	Items          *[]ListInstance `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceResponse", string(data)}, " ")
}
