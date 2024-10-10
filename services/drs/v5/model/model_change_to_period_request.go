package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeToPeriodRequest Request Object
type ChangeToPeriodRequest struct {

	// 转包周期的任务ID。
	JobId string `json:"job_id"`

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *ToPeriodReq `json:"body,omitempty"`
}

func (o ChangeToPeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeToPeriodRequest struct{}"
	}

	return strings.Join([]string{"ChangeToPeriodRequest", string(data)}, " ")
}
