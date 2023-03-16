package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAutoJobResponse struct {

	// 自动作业ID
	Id *string `json:"id,omitempty"`

	// 自动作业的名称
	Name *string `json:"name,omitempty"`

	// 自动作业的描述
	Description *string `json:"description,omitempty"`

	// 自动作业状态
	Status *string `json:"status,omitempty"`

	// 自动作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 自动作业结束时间
	FinishTime *string `json:"finish_time,omitempty"`

	// 自动作业依赖的数据库ID
	DatabaseId *string `json:"database_id,omitempty"`

	// 自动作业状态更新列
	DatabaseColumn *string `json:"database_column,omitempty"`

	// 自动作业状态更新列的类型，不填默认为EXISTED
	DatabaseColumnType *string `json:"database_column_type,omitempty"`

	// 是否清空作业状态更新列
	CleanDatabaseColumn *bool `json:"clean_database_column,omitempty"`

	// 自动作业触发器
	DatabaseTrigger *[]DatabaseTriggerDto `json:"database_trigger,omitempty"`

	ToolInfo *ToolInfoDto `json:"tool_info,omitempty"`

	// 作业的名称，取值范围：[1,63]，允许大小写字母、数字、以及特殊字符中划线(-)
	JobName *string `json:"job_name,omitempty"`

	// 作业的名称类型
	JobNameType *string `json:"job_name_type,omitempty"`

	// 作业的描述,取值范围：输入字符最大长度为255
	JobDescription *string `json:"job_description,omitempty"`

	// 作业标签，取值范围[0,5]，单个标签最大长度32字符，仅仅包含小写字母或数字或大写字母
	Labels *[]string `json:"labels,omitempty"`

	// 作业的优先级,取值范围[0,9]，0最低，默认数值0
	Priority *int32 `json:"priority,omitempty"`

	// 作业执行超时时长，取值范围: [1, 144000]，单位：分钟，默认数值1440
	Timeout *int32 `json:"timeout,omitempty"`

	// job结果存储目录，不指定则在workflow的工作目录下生产job同名子目录，指定则已指定路径为准;
	OutputDir *string `json:"output_dir,omitempty"`

	// 输出路径的类型
	OutputDirType *string `json:"output_dir_type,omitempty"`

	// 节点标签 取值范围[0,1]，单个标签最大长度63字符
	NodeLabels *[]string `json:"node_labels,omitempty"`

	// 自动作业使用的IO加速实例id，不填表示不使用
	IoAccId *string `json:"io_acc_id,omitempty"`

	// 自动作业依赖的流程信息
	Tasks          *[]TaskDefinitionDto `json:"tasks,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowAutoJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoJobResponse struct{}"
	}

	return strings.Join([]string{"ShowAutoJobResponse", string(data)}, " ")
}
