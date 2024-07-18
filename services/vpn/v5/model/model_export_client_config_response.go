package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportClientConfigResponse Response Object
type ExportClientConfigResponse struct {

	// 客户端配置
	ClientConfig *string `json:"client_config,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ExportClientConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportClientConfigResponse struct{}"
	}

	return strings.Join([]string{"ExportClientConfigResponse", string(data)}, " ")
}
