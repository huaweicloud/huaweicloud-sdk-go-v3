package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSearchFactoryBaselinesRequest Request Object
type ListSearchFactoryBaselinesRequest struct {

	// DataArts Studio实例ID。
	InstanceId string `json:"instance_id"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	WorkspaceId string `json:"workspace_id"`

	// 基线任务名称。
	Name *string `json:"name,omitempty"`

	// 责任人名称。
	OwnerName *string `json:"owner_name,omitempty"`

	// 优先级，取值有1/2/3/4/5，默认查询所有优先级。当同时查询优先级为1/2/3时，样例如下：priority=1&priority=2&priority=3。
	Priority *int32 `json:"priority,omitempty"`

	// 排序规则，取值如下： - priority asc: 按照优先级升序。 - priority desc: 按照优先级降序。 默认不排序。
	OrderBy *string `json:"order_by,omitempty"`

	// 开启基线任务。 true: 开启基线任务，系统将会监控基线任务以及其依赖链上游的所有任务。 false: 关闭基线任务，系统不会监控基线任务以及其依赖链上游的所有任务。 默认查询全部。
	Enable *bool `json:"enable,omitempty"`

	// 分页列表的页数，默认值为1。取值范围大于等于1。
	Offset *int32 `json:"offset,omitempty"`

	// 分页返回结果，指定每页最大记录数。范围[1,100] 默认值：10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListSearchFactoryBaselinesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSearchFactoryBaselinesRequest struct{}"
	}

	return strings.Join([]string{"ListSearchFactoryBaselinesRequest", string(data)}, " ")
}
