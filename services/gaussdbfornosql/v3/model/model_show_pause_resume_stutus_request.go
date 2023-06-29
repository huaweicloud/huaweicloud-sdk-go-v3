package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPauseResumeStutusRequest Request Object
type ShowPauseResumeStutusRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowPauseResumeStutusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPauseResumeStutusRequest struct{}"
	}

	return strings.Join([]string{"ShowPauseResumeStutusRequest", string(data)}, " ")
}
