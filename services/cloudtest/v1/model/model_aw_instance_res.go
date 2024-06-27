package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AwInstanceRes struct {

	// AW内容描述字段
	AliasRunaw *string `json:"alias_runaw,omitempty"`

	AuthenticationInfo *AuthInfo `json:"authentication_info,omitempty"`

	// 认证类型,0-无认证;1-aksk认证
	AuthenticationType *string `json:"authentication_type,omitempty"`

	// 脚本模板描述信息，在用例更新时添加
	AwDescription *string `json:"aw_description,omitempty"`

	// aw id
	AwId *string `json:"aw_id,omitempty"`

	// aw类型 0-setup;1-test;2-teardown
	AwType *int32 `json:"aw_type,omitempty"`

	BasicAw *BasicAwRes `json:"basic_aw,omitempty"`

	// instance的参数body体类型：form/text
	BodyParamType *string `json:"body_param_type,omitempty"`

	// 顺序
	ByOrder *int32 `json:"by_order,omitempty"`

	// change sign
	ChangeSign *int32 `json:"changeSign,omitempty"`

	// 检查点List
	CheckPointList *[]CheckPoint `json:"check_point_list,omitempty"`

	// aw实例
	Children *[]AwInstanceRes `json:"children,omitempty"`

	// 条件语句
	ConditionStatement *string `json:"condition_statement,omitempty"`

	// 条件类型 0-not condition;1-if begin;2-if
	ConditionType *int32 `json:"condition_type,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 创建时间戳
	CreateTimeStamp *int64 `json:"create_time_stamp,omitempty"`

	// 创建时间字符串
	CreateTimeString *string `json:"create_time_string,omitempty"`

	// 创建人
	CreateUser *string `json:"create_user,omitempty"`

	// 测试步骤自定义请求头List；后续自定义URL请求头不再使用该字段
	CustomHeader *[]AwParam `json:"custom_header,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	ErrorInfo *ErrorInfo `json:"error_info,omitempty"`

	ExtraInfo *ExtraInfo `json:"extra_info,omitempty"`

	// 该aw是否来自外部工程 0-no;1-yes
	FromOutside *int32 `json:"from_outside,omitempty"`

	// level
	HasLevel *int32 `json:"hasLevel,omitempty"`

	HeaderArray *[]ArrayNode `json:"header_array,omitempty"`

	// 值不为null表示老的IF判断语句；值为null表示新的IF判断语句
	HisScript *string `json:"his_script,omitempty"`

	// id
	Id *string `json:"id,omitempty"`

	// 是否模板类型测试步骤 0：自定义URL配置类型；1：模板类型测试步骤
	IsBasic *int32 `json:"is_basic,omitempty"`

	// 是否是契约AW 0-否；1-yes
	IsContractAw *int32 `json:"is_contract_aw,omitempty"`

	// 是否被禁用 0-否；1-yes
	IsDisabled *int32 `json:"is_disabled,omitempty"`

	// 是否是安全测试aw
	IsSectestAw *int32 `json:"is_sectest_aw,omitempty"`

	// 用例级别
	Level *int32 `json:"level,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 参数依赖规则
	ParamDependent *string `json:"param_dependent,omitempty"`

	// 是否启用参数依赖
	ParamDependentEnabled *int32 `json:"param_dependent_enabled,omitempty"`

	// 参数类型和参数值对应List
	ParamTypeAndValue *[]AwParam `json:"param_type_and_value,omitempty"`

	// awinstance所在的用例\\逻辑用例\\组合aw所属项目,该字段是新增字段,因此部分awinstance是无值的,所以只可写该值,而不能读取该值
	ProjectId *string `json:"projectId,omitempty"`

	// 区域名称
	Region *string `json:"region,omitempty"`

	// awId层级关系
	Relation *string `json:"relation,omitempty"`

	// relation id
	RelationId *string `json:"relation_id,omitempty"`

	// 映射类型 1-反向删除映射;2-用例自动添加的方向删除步骤
	RelationType *int32 `json:"relation_type,omitempty"`

	// 重试间隔时间 (ms) 为空表示不等待
	RetryInterval *string `json:"retry_interval,omitempty"`

	// 重试次数
	RetryTimes *string `json:"retry_times,omitempty"`

	// 获取脚本生成时，要使用的步骤名称
	ScriptName *string `json:"scriptName,omitempty"`

	// aw所来自工程的服务名和阶段名 fromOutside为1时该值有效
	ServiceAndStage *string `json:"service_and_stage,omitempty"`

	// 测试步骤来源
	SpecialType *int32 `json:"special_type,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 更新时间戳
	UpdateTimeStamp *int64 `json:"update_time_stamp,omitempty"`

	// 更新时间字符串
	UpdateTimeString *string `json:"update_time_string,omitempty"`

	// 更新人
	UpdateUser *string `json:"update_user,omitempty"`

	// user id
	UserId *string `json:"user_id,omitempty"`

	// 定义的变量信息
	VariableList *[]AwVariable `json:"variable_list,omitempty"`
}

func (o AwInstanceRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AwInstanceRes struct{}"
	}

	return strings.Join([]string{"AwInstanceRes", string(data)}, " ")
}
