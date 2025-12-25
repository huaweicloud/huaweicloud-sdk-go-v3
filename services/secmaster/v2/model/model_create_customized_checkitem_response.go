package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCustomizedCheckitemResponse Response Object
type CreateCustomizedCheckitemResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误描述
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCustomizedCheckitemResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCustomizedCheckitemResponse struct{}"
	}

	return strings.Join([]string{"CreateCustomizedCheckitemResponse", string(data)}, " ")
}
