package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactoryJobDependInstancesRequest Request Object
type ShowFactoryJobDependInstancesRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 作业名称。指定要查询上下游依赖关系的目标作业名称。 作业名称在创建作业时由用户指定，可通过\"查询作业列表\"接口获取。
	JobName string `json:"job_name"`

	// 查询的依赖方向。用于指定查询作业的直接上游、直接下游或同时查询上下游关系。 取值范围： - parent：查询直接上游作业，即当前作业依赖的作业。 - child：查询直接下游作业，即依赖当前作业的作业。 - both：同时查询直接上游和直接下游作业。 默认取值：both
	Relation *string `json:"relation,omitempty"`
}

func (o ShowFactoryJobDependInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryJobDependInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowFactoryJobDependInstancesRequest", string(data)}, " ")
}
