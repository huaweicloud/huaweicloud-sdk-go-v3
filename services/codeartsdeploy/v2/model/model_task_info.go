package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskInfo 查询应用详情响应体
type TaskInfo struct {

	// 部署任务id
	TaskId *string `json:"task_id,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 部署类型模式，包括deployTemplate，ansible，shell
	DeploySystem *string `json:"deploy_system,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 应用状态，Draft表示草稿状态，Available表示可用状态
	State *TaskInfoState `json:"state,omitempty"`

	// 最后一次部署时间
	ExecutionTime *string `json:"execution_time,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否使用默认权限矩阵
	IsDefautPermission *bool `json:"is_defaut_permission,omitempty"`

	// 模板id
	TemplateId *string `json:"template_id,omitempty"`

	// 应用创建者用户名
	Owner *string `json:"owner,omitempty"`

	// 应用创建者昵称
	NickName *string `json:"nick_name,omitempty"`

	// 应用创建者用户ID
	OwnerId *string `json:"owner_id,omitempty"`

	// 应用创建者租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 应用创建者租户名
	TenantName *string `json:"tenant_name,omitempty"`

	// slave集群id，默认为null时使用默认slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty"`

	// 当前用户是否已收藏
	IsCare *bool `json:"is_care,omitempty"`

	// 是否有编辑权限
	CanModify *bool `json:"can_modify,omitempty"`

	// 是否有删除的权限
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty"`

	// 是否有部署权限
	CanExecute *bool `json:"can_execute,omitempty"`

	// 是否有复制权限
	CanCopy *bool `json:"can_copy,omitempty"`

	// 是否有管理权限，包含增删改查部署以及权限修改
	CanManage *bool `json:"can_manage,omitempty"`

	// 应用和AOM应用组件对应关系
	AppComponentList *[]AppComponentDao `json:"app_component_list,omitempty"`

	// 角色ID,0：应用创建者，-1：项目创建者，3：项目经理，4：开发人员，5：测试经理，6：测试人员，7：参与者，8：浏览者
	RoleId *int32 `json:"role_id,omitempty"`

	// 部署任务id
	Id *string `json:"id,omitempty"`

	// 部署记录序列号
	ReleaseId *int32 `json:"release_id,omitempty"`

	// 部署时间
	Duration *string `json:"duration,omitempty"`

	// 部署状态
	ExecutionState *string `json:"execution_state,omitempty"`

	// 部署者id
	ExecutorId *string `json:"executor_id,omitempty"`

	// 部署者名称
	ExecutorNickName *string `json:"executor_nick_name,omitempty"`

	// 部署步骤
	Steps map[string]Step `json:"steps,omitempty"`
}

func (o TaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskInfo struct{}"
	}

	return strings.Join([]string{"TaskInfo", string(data)}, " ")
}

type TaskInfoState struct {
	value string
}

type TaskInfoStateEnum struct {
	AVAILABLE TaskInfoState
	DRAFT     TaskInfoState
}

func GetTaskInfoStateEnum() TaskInfoStateEnum {
	return TaskInfoStateEnum{
		AVAILABLE: TaskInfoState{
			value: "Available",
		},
		DRAFT: TaskInfoState{
			value: "Draft",
		},
	}
}

func (c TaskInfoState) Value() string {
	return c.value
}

func (c TaskInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskInfoState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
