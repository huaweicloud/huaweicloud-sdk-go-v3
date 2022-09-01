package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDeploymentsResponse struct {

	// 应用部署总数
	Count *int64 `json:"count,omitempty" xml:"count"`

	// 应用部署列表
	Deployments    *[]DeploymentResp `json:"deployments,omitempty" xml:"deployments"`
	HttpStatusCode int               `json:"-"`
}

func (o ListDeploymentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeploymentsResponse struct{}"
	}

	return strings.Join([]string{"ListDeploymentsResponse", string(data)}, " ")
}
