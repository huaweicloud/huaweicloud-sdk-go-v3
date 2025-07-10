package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentsInstallConditionResponse Response Object
type ListAgentsInstallConditionResponse struct {

	// 桌面agent安装情况。
	AgentsCondition *[]AgentsCondition `json:"agents_condition,omitempty"`

	// 总共条数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAgentsInstallConditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentsInstallConditionResponse struct{}"
	}

	return strings.Join([]string{"ListAgentsInstallConditionResponse", string(data)}, " ")
}
