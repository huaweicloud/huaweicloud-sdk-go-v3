package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmApplicationTypesResponse Response Object
type ConfirmApplicationTypesResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 类型名称
	Items          *[]string `json:"items,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ConfirmApplicationTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmApplicationTypesResponse struct{}"
	}

	return strings.Join([]string{"ConfirmApplicationTypesResponse", string(data)}, " ")
}
