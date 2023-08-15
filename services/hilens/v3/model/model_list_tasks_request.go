package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksRequest Request Object
type ListTasksRequest struct {

	// 部署ID，从专业版HiLens控制台部署管理[获取部署列表](getDeploymentListUsingGET.xml)获取
	DeploymentId string `json:"deployment_id"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量，取值范围1~100，默认为100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksRequest struct{}"
	}

	return strings.Join([]string{"ListTasksRequest", string(data)}, " ")
}
