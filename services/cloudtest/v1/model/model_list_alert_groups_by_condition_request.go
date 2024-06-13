package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlertGroupsByConditionRequest Request Object
type ListAlertGroupsByConditionRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *ListAlertGroupsByConditionRequestBody `json:"body,omitempty"`
}

func (o ListAlertGroupsByConditionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlertGroupsByConditionRequest struct{}"
	}

	return strings.Join([]string{"ListAlertGroupsByConditionRequest", string(data)}, " ")
}
