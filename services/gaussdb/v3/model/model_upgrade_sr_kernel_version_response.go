package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeSrKernelVersionResponse Response Object
type UpgradeSrKernelVersionResponse struct {

	// **参数解释**： 升级工作流ID。  **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpgradeSrKernelVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeSrKernelVersionResponse struct{}"
	}

	return strings.Join([]string{"UpgradeSrKernelVersionResponse", string(data)}, " ")
}
