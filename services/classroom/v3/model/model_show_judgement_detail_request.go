package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowJudgementDetailRequest struct {
	// 判题任务ID

	JudgementId string `json:"judgement_id"`
}

func (o ShowJudgementDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJudgementDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowJudgementDetailRequest", string(data)}, " ")
}
