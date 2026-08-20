package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDeploymentSolutionResponse Response Object
type ChangeDeploymentSolutionResponse struct {

	// **参数解释**: 形态变更的任务ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeDeploymentSolutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDeploymentSolutionResponse struct{}"
	}

	return strings.Join([]string{"ChangeDeploymentSolutionResponse", string(data)}, " ")
}
