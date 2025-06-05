package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKeystorePermissionResponse Response Object
type ShowKeystorePermissionResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *string `json:"error,omitempty"`

	Result         *ShowKeystorePermissionResponseBodyResult `json:"result,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ShowKeystorePermissionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeystorePermissionResponse struct{}"
	}

	return strings.Join([]string{"ShowKeystorePermissionResponse", string(data)}, " ")
}
