package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppDisableStatusRequestBody 禁用/取消禁用应用请求体
type UpdateAppDisableStatusRequestBody struct {

	// 是否禁用， true禁用， false取消禁用
	IsDisable bool `json:"is_disable"`
}

func (o UpdateAppDisableStatusRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppDisableStatusRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAppDisableStatusRequestBody", string(data)}, " ")
}
