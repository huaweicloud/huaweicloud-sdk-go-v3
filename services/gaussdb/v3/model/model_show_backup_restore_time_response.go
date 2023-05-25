package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBackupRestoreTimeResponse struct {

	// 可恢复时间段列表。
	RestoreTimes   *[]RestoreTimeInfo `json:"restore_times,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowBackupRestoreTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupRestoreTimeResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupRestoreTimeResponse", string(data)}, " ")
}
