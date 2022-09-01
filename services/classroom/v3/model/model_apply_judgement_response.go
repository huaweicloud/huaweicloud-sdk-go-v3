package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ApplyJudgementResponse struct {

	// 判题任务ID
	JudgementId    *string `json:"judgement_id,omitempty" xml:"judgement_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ApplyJudgementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyJudgementResponse struct{}"
	}

	return strings.Join([]string{"ApplyJudgementResponse", string(data)}, " ")
}
