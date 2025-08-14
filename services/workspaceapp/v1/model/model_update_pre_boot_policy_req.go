package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePreBootPolicyReq 批量设置应用预启动请求体。
type UpdatePreBootPolicyReq struct {

	// 应用ID,最多同时操作5个。
	Ids []string `json:"ids"`

	// 是否设置应用预启动。
	IsPreBoot bool `json:"is_pre_boot"`
}

func (o UpdatePreBootPolicyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePreBootPolicyReq struct{}"
	}

	return strings.Join([]string{"UpdatePreBootPolicyReq", string(data)}, " ")
}
