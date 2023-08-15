package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDeploymentsResponse Response Object
type ListDeploymentsResponse struct {

	// 应用部署总数
	Count *int64 `json:"count,omitempty"`

	// 应用部署列表
	Deployments    *[]DeploymentResp `json:"deployments,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListDeploymentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeploymentsResponse struct{}"
	}

	return strings.Join([]string{"ListDeploymentsResponse", string(data)}, " ")
}
