package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubTestCaseByConditionsUsingRequest Request Object
type ListSubTestCaseByConditionsUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *SubTaskCaseQuery `json:"body,omitempty"`
}

func (o ListSubTestCaseByConditionsUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubTestCaseByConditionsUsingRequest struct{}"
	}

	return strings.Join([]string{"ListSubTestCaseByConditionsUsingRequest", string(data)}, " ")
}
