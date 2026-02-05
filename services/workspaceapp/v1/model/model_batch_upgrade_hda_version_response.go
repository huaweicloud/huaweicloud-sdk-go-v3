package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpgradeHdaVersionResponse Response Object
type BatchUpgradeHdaVersionResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchUpgradeHdaVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpgradeHdaVersionResponse struct{}"
	}

	return strings.Join([]string{"BatchUpgradeHdaVersionResponse", string(data)}, " ")
}
