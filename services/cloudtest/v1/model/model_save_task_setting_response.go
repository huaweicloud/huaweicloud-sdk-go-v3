package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveTaskSettingResponse Response Object
type SaveTaskSettingResponse struct {

	// 错误编码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误原因
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SaveTaskSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveTaskSettingResponse struct{}"
	}

	return strings.Join([]string{"SaveTaskSettingResponse", string(data)}, " ")
}
