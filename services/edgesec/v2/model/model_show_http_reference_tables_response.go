package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpReferenceTablesResponse Response Object
type ShowHttpReferenceTablesResponse struct {

	// 引用表数量
	Total *int32 `json:"total,omitempty"`

	// 引用表列表
	Items          *[]ShowHttpReferenceTableResponseBody `json:"items,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ShowHttpReferenceTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpReferenceTablesResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpReferenceTablesResponse", string(data)}, " ")
}
