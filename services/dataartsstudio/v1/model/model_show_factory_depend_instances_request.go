package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactoryDependInstancesRequest Request Object
type ShowFactoryDependInstancesRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 作业实例id，可通过作业实例列表接口获取。
	InstanceId int64 `json:"instance_id"`

	// 支持选择查询实例的直接上游、直接下游或者是直接上下游，取值为 parent、child、both，默认为both。 - parent：直接上游实例 - child：直接下游实例 - both：直接上下游实例
	Relation *string `json:"relation,omitempty"`

	// 默认值为1, depth是上下游依赖查询的层级深度，例如 depth=1 只查直接依赖，depth=2 查到依赖的依赖。单次查询可返回的最大深度为50层。
	Depth *int32 `json:"depth,omitempty"`

	// 默认值为true, 当latest=true的时候，控制是否只返回每个依赖任务的最新实例，true 时只返回endTime 最晚的执行记录，false 时返回所有历史实例。
	Latest *bool `json:"latest,omitempty"`
}

func (o ShowFactoryDependInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactoryDependInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowFactoryDependInstancesRequest", string(data)}, " ")
}
