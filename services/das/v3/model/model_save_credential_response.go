package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveCredentialResponse Response Object
type SaveCredentialResponse struct {

	// 保存AK/SK是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o SaveCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveCredentialResponse struct{}"
	}

	return strings.Join([]string{"SaveCredentialResponse", string(data)}, " ")
}
