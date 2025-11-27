package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStatisticsRequest Request Object
type UpdateStatisticsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	Body *UpdateStatisticsRequestBody `json:"body,omitempty"`
}

func (o UpdateStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStatisticsRequest struct{}"
	}

	return strings.Join([]string{"UpdateStatisticsRequest", string(data)}, " ")
}
