package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiAppSettingRequest AI应用配置请求体
type AiAppSettingRequest struct {
	ApiSign *ApiSign `json:"api_sign,omitempty"`
}

func (o AiAppSettingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiAppSettingRequest struct{}"
	}

	return strings.Join([]string{"AiAppSettingRequest", string(data)}, " ")
}
