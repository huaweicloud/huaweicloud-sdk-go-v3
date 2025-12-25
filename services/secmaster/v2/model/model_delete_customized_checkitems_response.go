package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCustomizedCheckitemsResponse Response Object
type DeleteCustomizedCheckitemsResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data           *BatchOperateBaselineResult `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o DeleteCustomizedCheckitemsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomizedCheckitemsResponse struct{}"
	}

	return strings.Join([]string{"DeleteCustomizedCheckitemsResponse", string(data)}, " ")
}
