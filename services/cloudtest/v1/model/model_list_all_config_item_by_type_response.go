package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllConfigItemByTypeResponse Response Object
type ListAllConfigItemByTypeResponse struct {

	// 错误编码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误原因
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAllConfigItemByTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllConfigItemByTypeResponse struct{}"
	}

	return strings.Join([]string{"ListAllConfigItemByTypeResponse", string(data)}, " ")
}
