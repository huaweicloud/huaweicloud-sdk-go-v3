package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddKeystorePermissionResponse Response Object
type AddKeystorePermissionResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	// 结果
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AddKeystorePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddKeystorePermissionResponse struct{}"
	}

	return strings.Join([]string{"AddKeystorePermissionResponse", string(data)}, " ")
}
