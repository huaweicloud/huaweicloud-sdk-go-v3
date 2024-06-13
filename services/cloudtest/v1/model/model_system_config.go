package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SystemConfig struct {

	// 配置项主键
	Id *string `json:"id,omitempty"`

	// 系统配置名称
	Key *string `json:"key,omitempty"`

	// 系统配置状态
	Value *string `json:"value,omitempty"`

	// 描述
	Remark *string `json:"remark,omitempty"`

	// region_id
	RegionId *string `json:"region_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 更新人名称
	UpdateName *string `json:"update_name,omitempty"`

	// 更新人编号
	UpdateNum *string `json:"update_num,omitempty"`
}

func (o SystemConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemConfig struct{}"
	}

	return strings.Join([]string{"SystemConfig", string(data)}, " ")
}
