package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCheckitemResponse Response Object
type UpdateCheckitemResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误描述
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateCheckitemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCheckitemResponse struct{}"
	}

	return strings.Join([]string{"UpdateCheckitemResponse", string(data)}, " ")
}
