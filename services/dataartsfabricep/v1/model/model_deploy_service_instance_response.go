package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployServiceInstanceResponse Response Object
type DeployServiceInstanceResponse struct {

	// 实例Id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeployServiceInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployServiceInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeployServiceInstanceResponse", string(data)}, " ")
}
