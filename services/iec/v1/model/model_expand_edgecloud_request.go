package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExpandEdgecloudRequest struct {
	// 部署计划ID。  约束： - 该接口只能执行指定边缘业务ID（id）创建的部署计划。

	DeploymentId string `json:"deployment_id"`
}

func (o ExpandEdgecloudRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExpandEdgecloudRequest struct{}"
	}

	return strings.Join([]string{"ExpandEdgecloudRequest", string(data)}, " ")
}
