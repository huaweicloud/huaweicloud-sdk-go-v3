package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeInstanceMinorVersionResponse Response Object
type UpgradeInstanceMinorVersionResponse struct {

	// 后台任务ID。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeInstanceMinorVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeInstanceMinorVersionResponse struct{}"
	}

	return strings.Join([]string{"UpgradeInstanceMinorVersionResponse", string(data)}, " ")
}
