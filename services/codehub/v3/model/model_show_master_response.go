package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMasterResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	// 响应结果
	Result *bool `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMasterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMasterResponse struct{}"
	}

	return strings.Join([]string{"ShowMasterResponse", string(data)}, " ")
}
