package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDeploymentJobsResponse struct {

	// 部署任务编号
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 任务ID
	JobId *string `json:"job_id,omitempty" xml:"job_id"`

	// 部署状态,-2：环境准备未就绪 -1 资源准备就绪 0 部署中 1：成功 2：失败
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 访问地址
	Address *string `json:"address,omitempty" xml:"address"`

	// 部署参数
	DeployParameters *string `json:"deploy_parameters,omitempty" xml:"deploy_parameters"`

	// 部署耗时
	Time *int32 `json:"time,omitempty" xml:"time"`

	// 创建人
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name"`

	// 创建时间
	CreatedTime    *string `json:"created_time,omitempty" xml:"created_time"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDeploymentJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDeploymentJobsResponse struct{}"
	}

	return strings.Join([]string{"CreateDeploymentJobsResponse", string(data)}, " ")
}
