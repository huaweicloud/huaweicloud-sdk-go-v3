package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建保护组请求体
type CreateProtectionGroupRequestBody struct {
	ServerGroup *CreateProtectionGroupRequestParams `json:"server_group"`
}

func (o CreateProtectionGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProtectionGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProtectionGroupRequestBody", string(data)}, " ")
}
