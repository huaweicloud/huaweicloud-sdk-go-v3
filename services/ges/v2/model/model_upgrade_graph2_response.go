package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeGraph2Response Response Object
type UpgradeGraph2Response struct {

	// 执行该异步任务的jobId。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeGraph2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeGraph2Response struct{}"
	}

	return strings.Join([]string{"UpgradeGraph2Response", string(data)}, " ")
}
