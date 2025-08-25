package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppAccessKeyResponse Response Object
type CreateAppAccessKeyResponse struct {

	// 访问密钥的id
	AccessKeyId    *string `json:"access_key_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppAccessKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"CreateAppAccessKeyResponse", string(data)}, " ")
}
