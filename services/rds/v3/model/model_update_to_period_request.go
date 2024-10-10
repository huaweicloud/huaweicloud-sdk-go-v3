package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateToPeriodRequest Request Object
type UpdateToPeriodRequest struct {

	// 操作的实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage string `json:"X-Language"`

	Body *ToPeriodReq `json:"body,omitempty"`
}

func (o UpdateToPeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateToPeriodRequest struct{}"
	}

	return strings.Join([]string{"UpdateToPeriodRequest", string(data)}, " ")
}
