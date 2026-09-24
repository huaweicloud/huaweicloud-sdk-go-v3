package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpExtendInfoUpdateExpirationTime struct {

	// 本次任务受影响的备份个数
	AffectedBackupsCount *int32 `json:"affected_backups_count,omitempty"`

	// 本次任务预期过期日期，格式：YYYY-MM-DD。
	ExpirationDay *string `json:"expiration_day,omitempty"`
}

func (o OpExtendInfoUpdateExpirationTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpExtendInfoUpdateExpirationTime struct{}"
	}

	return strings.Join([]string{"OpExtendInfoUpdateExpirationTime", string(data)}, " ")
}
