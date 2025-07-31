package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListK8sDeploymentsResponse Response Object
type ListK8sDeploymentsResponse struct {

	// deployment实例总数
	TotalNum *int64 `json:"total_num,omitempty"`

	// 最近更新时间
	LastUpdateTime *int64 `json:"last_update_time,omitempty"`

	// 资源类型，包含如下四种。   - deploy ：无状态负载   - sts ：有状态负载   - job ： 普通任务   - cronjob ：定时任务
	Type *string `json:"type,omitempty"`

	// resources基本信息列表
	ResourcesInfoList *[]ServerlessDeploymentInfo `json:"resources_info_list,omitempty"`
	HttpStatusCode    int                         `json:"-"`
}

func (o ListK8sDeploymentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListK8sDeploymentsResponse struct{}"
	}

	return strings.Join([]string{"ListK8sDeploymentsResponse", string(data)}, " ")
}
