package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessKeyResponse Response Object
type ShowAccessKeyResponse struct {

	// 访问密钥的access key
	AccessKey *string `json:"access_key,omitempty"`

	// 访问密钥的secret key
	SecretKey *string `json:"secret_key,omitempty"`

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	// 密钥名称
	KeyName *string `json:"key_name,omitempty"`

	// 是否导入
	IsImported     *bool `json:"is_imported,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowAccessKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessKeyResponse", string(data)}, " ")
}
