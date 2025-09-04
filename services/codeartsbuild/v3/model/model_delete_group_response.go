package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupResponse Response Object
type DeleteGroupResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteGroupResponse", string(data)}, " ")
}
