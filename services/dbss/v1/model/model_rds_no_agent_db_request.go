package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RdsNoAgentDbRequest struct {

	// 添加数据库信息列表
	Databases []RdsNoAgentDbRequestDatabases `json:"databases"`

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`
}

func (o RdsNoAgentDbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RdsNoAgentDbRequest struct{}"
	}

	return strings.Join([]string{"RdsNoAgentDbRequest", string(data)}, " ")
}
