package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUpdateBackupEnhancePolicyResponse Response Object
type ListUpdateBackupEnhancePolicyResponse struct {

	// 备份时间段开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 备份时间段结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 手动备份保留时长
	ManualBackupRetentionDays *string `json:"manual_backup_retention_days,omitempty"`

	// 高频备份的频率
	Frequency *string `json:"frequency,omitempty"`

	// 备份策略集
	Policies       *[]Policy `json:"policies,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListUpdateBackupEnhancePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUpdateBackupEnhancePolicyResponse struct{}"
	}

	return strings.Join([]string{"ListUpdateBackupEnhancePolicyResponse", string(data)}, " ")
}
