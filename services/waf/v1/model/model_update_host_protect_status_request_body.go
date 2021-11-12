package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改域名防护状态请求体
type UpdateHostProtectStatusRequestBody struct {
	// 防护状态（-1：bypass直接放行，0：暂停防护，1：开启防护）

	ProtectStatus *int32 `json:"protect_status,omitempty"`
}

func (o UpdateHostProtectStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostProtectStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHostProtectStatusRequestBody", string(data)}, " ")
}
