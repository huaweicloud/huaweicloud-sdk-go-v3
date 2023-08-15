package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PauseResumeDataSynchronizationResponse Response Object
type PauseResumeDataSynchronizationResponse struct {

	// 暂停/恢复具备容灾关系的实例数据同步的工作ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PauseResumeDataSynchronizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseResumeDataSynchronizationResponse struct{}"
	}

	return strings.Join([]string{"PauseResumeDataSynchronizationResponse", string(data)}, " ")
}
