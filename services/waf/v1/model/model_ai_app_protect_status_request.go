package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiAppProtectStatusRequest struct {

	// 防护状态（0：禁用，1：启用）
	ProtectStatus int32 `json:"protect_status"`
}

func (o AiAppProtectStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiAppProtectStatusRequest struct{}"
	}

	return strings.Join([]string{"AiAppProtectStatusRequest", string(data)}, " ")
}
