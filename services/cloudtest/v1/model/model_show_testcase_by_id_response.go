package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTestcaseByIdResponse Response Object
type ShowTestcaseByIdResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	Data *TestCase `json:"data,omitempty"`

	// 错误信息
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTestcaseByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTestcaseByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowTestcaseByIdResponse", string(data)}, " ")
}
