package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmProtectionTypesResponse Response Object
type ConfirmProtectionTypesResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 类型名称
	Items          *[]string `json:"items,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ConfirmProtectionTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmProtectionTypesResponse struct{}"
	}

	return strings.Join([]string{"ConfirmProtectionTypesResponse", string(data)}, " ")
}
