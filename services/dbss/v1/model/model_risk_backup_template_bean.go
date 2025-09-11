package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RiskBackupTemplateBean struct {

	// 数据库ID
	DbId *string `json:"db_id,omitempty"`

	// 数据库IP
	DbIp *string `json:"db_ip,omitempty"`

	// 数据库名称
	DbName *string `json:"db_name,omitempty"`

	// 数据库端口
	DbPort *int32 `json:"db_port,omitempty"`

	// 配置ID
	Id *string `json:"id,omitempty"`

	// 状态 - 0: 关闭 - 1：开启
	Status *int32 `json:"status,omitempty"`
}

func (o RiskBackupTemplateBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskBackupTemplateBean struct{}"
	}

	return strings.Join([]string{"RiskBackupTemplateBean", string(data)}, " ")
}
