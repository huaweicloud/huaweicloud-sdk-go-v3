package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListDeploymentsResponse struct {
	// 部署计划列表的总和。

	Count *int32 `json:"count,omitempty"`
	// 部署计划列表。

	Deployments    *[]Deployment `json:"deployments,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDeploymentsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDeploymentsResponse struct{}"
	}

	return strings.Join([]string{"ListDeploymentsResponse", string(data)}, " ")
}
