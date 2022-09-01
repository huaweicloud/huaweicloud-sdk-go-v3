package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDeploymentsResponse struct {

	// 部署计划列表的总和。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 部署计划列表。
	Deployments    *[]Deployment `json:"deployments,omitempty" xml:"deployments"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDeploymentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDeploymentsResponse struct{}"
	}

	return strings.Join([]string{"ListDeploymentsResponse", string(data)}, " ")
}
