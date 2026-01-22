package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreRecycleInstanceResponse Response Object
type RestoreRecycleInstanceResponse struct {

	// **参数解释**： 实例恢复列表。
	Results        *[]InstanceResumeResult `json:"results,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o RestoreRecycleInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreRecycleInstanceResponse struct{}"
	}

	return strings.Join([]string{"RestoreRecycleInstanceResponse", string(data)}, " ")
}
