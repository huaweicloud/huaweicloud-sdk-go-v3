package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RiskBackupTemplate struct {

	// 备份周期 - PER_DAY：每天 - PER_WEEK：每周 - PER_MONTH：每月 - PER_HOUR：每小时 - FIVE_MIN：每5分钟
	Cycle *string `json:"cycle,omitempty"`

	// 数据库ID
	DbId string `json:"db_id"`

	// 备份开关    - 0：关闭    - 1：开启
	Status int32 `json:"status"`
}

func (o RiskBackupTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RiskBackupTemplate struct{}"
	}

	return strings.Join([]string{"RiskBackupTemplate", string(data)}, " ")
}
