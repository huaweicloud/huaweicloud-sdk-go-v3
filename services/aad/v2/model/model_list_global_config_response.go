package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGlobalConfigResponse Response Object
type ListGlobalConfigResponse struct {

	// 支持的tls版本列表
	SupportTls *[]string `json:"support_tls,omitempty"`

	// 加密套件列表
	Cipher         *[]CipherInfo `json:"cipher,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListGlobalConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGlobalConfigResponse struct{}"
	}

	return strings.Join([]string{"ListGlobalConfigResponse", string(data)}, " ")
}
