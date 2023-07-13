package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAlarmTemplatesRequest Request Object
type BatchDeleteAlarmTemplatesRequest struct {
	Body *BatchDeleteAlarmTemplatesRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteAlarmTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlarmTemplatesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlarmTemplatesRequest", string(data)}, " ")
}
