package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBigkeyScanTaskDetailsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 大key分析任务ID。
	BigkeyId string `json:"bigkey_id" xml:"bigkey_id"`
}

func (o ShowBigkeyScanTaskDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBigkeyScanTaskDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowBigkeyScanTaskDetailsRequest", string(data)}, " ")
}
