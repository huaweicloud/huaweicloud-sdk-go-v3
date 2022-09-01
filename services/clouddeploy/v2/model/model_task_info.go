package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询部署任务详情响应体
type TaskInfo struct {

	// 部署任务id
	TaskId *string `json:"task_id,omitempty" xml:"task_id"`

	// 部署任务名称
	Name *string `json:"name,omitempty" xml:"name"`

	// devcloud项目id
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// devcloud项目名称
	ProjectName *string `json:"project_name,omitempty" xml:"project_name"`

	// 部署类型模式，包括deployTemplate，ansible，shell
	DeploySystem *string `json:"deploy_system,omitempty" xml:"deploy_system"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty" xml:"create_time"`

	// 修改时间
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`

	// 部署任务状态，Draft表示草稿状态，Available表示可用状态
	State *TaskInfoState `json:"state,omitempty" xml:"state"`

	// 最后一次执行时间
	ExecutionTime *string `json:"execution_time,omitempty" xml:"execution_time"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 是否使用默认权限矩阵
	IsDefautPermission *bool `json:"is_defaut_permission,omitempty" xml:"is_defaut_permission"`

	// 模板id
	TemplateId *string `json:"template_id,omitempty" xml:"template_id"`

	// 部署任务创建者用户名
	Owner *string `json:"owner,omitempty" xml:"owner"`

	// 部署任务创建者昵称
	NickName *string `json:"nick_name,omitempty" xml:"nick_name"`

	// 部署任务创建者用户ID
	OwnerId *string `json:"owner_id,omitempty" xml:"owner_id"`

	// 部署任务创建者租户ID
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id"`

	// 部署任务创建者租户名
	TenantName *string `json:"tenant_name,omitempty" xml:"tenant_name"`

	// slave集群id，默认为null时使用devcloud八爪鱼slave集群，用户自定义slave时为slave集群id
	SlaveClusterId *string `json:"slave_cluster_id,omitempty" xml:"slave_cluster_id"`

	// 当前用户是否已收藏
	IsCare *bool `json:"is_care,omitempty" xml:"is_care"`

	// 是否有编辑权限
	CanModify *bool `json:"can_modify,omitempty" xml:"can_modify"`

	// 是否有删除的权限
	CanDelete *bool `json:"can_delete,omitempty" xml:"can_delete"`

	// 是否有查看权限
	CanView *bool `json:"can_view,omitempty" xml:"can_view"`

	// 是否有执行权限
	CanExecute *bool `json:"can_execute,omitempty" xml:"can_execute"`

	// 是否有复制权限
	CanCopy *bool `json:"can_copy,omitempty" xml:"can_copy"`

	// 是否有管理权限，包含增删改查执行以及权限修改
	CanManage *bool `json:"can_manage,omitempty" xml:"can_manage"`

	// 部署任务和应用组件对应关系
	AppComponentList *[]AppComponentDao `json:"app_component_list,omitempty" xml:"app_component_list"`

	// 角色ID,0：部署任务创建者，-1：项目创建者，3：项目经理，4：开发人员，5：测试经理，6：测试人员，7：参与者，8：浏览者
	RoleId *int32 `json:"role_id,omitempty" xml:"role_id"`
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
