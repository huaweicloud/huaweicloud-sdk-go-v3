package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDeploymentsResponse struct {

	// 部署的全部数量
	Count *int32 `json:"count,omitempty"`

	// 部署列表信息
	Deployments    *[]Deployment `json:"deployments,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowDeploymentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentsResponse struct{}"
	}

	return strings.Join([]string{"ShowDeploymentsResponse", string(data)}, " ")
}
