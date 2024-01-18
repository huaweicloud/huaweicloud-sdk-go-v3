package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTsvi 服务器更新虚拟会话IP配置请求。
type UpdateTsvi struct {

	// 服务器ID。
	Id string `json:"id"`

	// **⚠ 预留字段，不使用，是否启用虚拟IP功能与服务器组配置保持一致。 是否启用虚拟IP功能。 开关只在租户配置允许启用虚拟IP场景有效，否则忽略传值并设置为关闭。
	Enable bool `json:"enable"`
}

func (o UpdateTsvi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTsvi struct{}"
	}

	return strings.Join([]string{"UpdateTsvi", string(data)}, " ")
}
