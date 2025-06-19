package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKeystoreResponse Response Object
type DeleteKeystoreResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteKeystoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKeystoreResponse struct{}"
	}

	return strings.Join([]string{"DeleteKeystoreResponse", string(data)}, " ")
}
