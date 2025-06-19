package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTheJobResponse Response Object
type DeleteTheJobResponse struct {
	Result *DeleteTheJobResponseBodyResult `json:"result,omitempty"`

	// 状态信息
	Status *string `json:"status,omitempty"`

	// 返回错误信息
	Error          *string `json:"error,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteTheJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTheJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteTheJobResponse", string(data)}, " ")
}
