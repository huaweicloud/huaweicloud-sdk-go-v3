package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTestsuiteInfoUsingRequest Request Object
type UpdateTestsuiteInfoUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	// 任务id
	SuiteId string `json:"suite_id"`

	Body *TaskInfoV4VoReq `json:"body,omitempty"`
}

func (o UpdateTestsuiteInfoUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestsuiteInfoUsingRequest struct{}"
	}

	return strings.Join([]string{"UpdateTestsuiteInfoUsingRequest", string(data)}, " ")
}
