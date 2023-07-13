package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConfigSettingResponse Response Object
type ShowConfigSettingResponse struct {

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 迁移类型
	MigrateType *string `json:"migrate_type,omitempty"`

	// 配置项的具体配置信息
	Configurations *string `json:"configurations,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowConfigSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConfigSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowConfigSettingResponse", string(data)}, " ")
}
