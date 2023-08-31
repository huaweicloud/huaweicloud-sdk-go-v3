package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationResponse Response Object
type DeleteApplicationResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	Result         *AppBaseResponse `json:"result,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o DeleteApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationResponse struct{}"
	}

	return strings.Join([]string{"DeleteApplicationResponse", string(data)}, " ")
}
