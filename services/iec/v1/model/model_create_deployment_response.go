package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateDeploymentResponse struct {
	// 部署计划ID。

	Id *string `json:"id,omitempty"`
	// 部署位置信息列表。

	Locations      *[]Location `json:"locations,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateDeploymentResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateDeploymentResponse struct{}"
	}

	return strings.Join([]string{"CreateDeploymentResponse", string(data)}, " ")
}
