package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateTestCaseListInfo struct {

	// 对外测试步骤
	Steps *[]TestCaseStepInfo `json:"steps,omitempty"`

	// 对外附件
	Attachments *[]TestCaseAttachmentInfo `json:"attachments,omitempty"`

	// 创建人
	Author *string `json:"author,omitempty"`

	// 名称
	Name string `json:"name"`

	// 用例等级
	Rank *int32 `json:"rank,omitempty"`

	// 处理人
	Owner *string `json:"owner,omitempty"`

	// 前置条件
	Preparation *string `json:"preparation,omitempty"`

	// 备注
	Remark *string `json:"remark,omitempty"`

	// 测试阶段
	Stage *string `json:"stage,omitempty"`

	// 测试类型
	Activity *string `json:"activity,omitempty"`

	// 关键词
	Keywords *string `json:"keywords,omitempty"`

	// apitest标记是否代码已提交
	Market *string `json:"market,omitempty"`

	// 设计者
	Designer *string `json:"designer,omitempty"`

	// 标签
	Tags *string `json:"tags,omitempty"`

	// 逻辑region，外部使用公有云实际区域，内部使用默认值
	Region *string `json:"region,omitempty"`

	// 对外关联资源类型
	RelateType *string `json:"relate_type,omitempty"`

	// 服务类型
	ServiceType *int32 `json:"service_type,omitempty"`

	// 对外只更新接口用例的java脚本路径标识
	OnlyChangeScript *string `json:"only_change_script,omitempty"`

	// 对外需求添加到迭代标识
	AddToIterator *string `json:"add_to_iterator,omitempty"`

	// 是否修改关联关系
	NeedUpdateRelation *string `json:"need_update_relation,omitempty"`

	// 创建版本Uri
	CreationVersionUri *string `json:"creation_version_uri,omitempty"`

	// 用例编号
	Number *string `json:"number,omitempty"`

	// 用例类型
	CaseType *int32 `json:"case_type,omitempty"`

	// 执行平台类型
	PlatformType *int32 `json:"platform_type,omitempty"`

	// 测试类型
	TestType *int32 `json:"test_type,omitempty"`

	// 设计描述
	DesignNote *string `json:"design_note,omitempty"`

	// 测试步骤
	TestStep *string `json:"test_step,omitempty"`

	// 期望结果
	ExpectOutput *string `json:"expect_output,omitempty"`

	// 测试环境类型
	EnvType *string `json:"env_type,omitempty"`

	// 执行平台
	ExePlatform *string `json:"exe_platform,omitempty"`

	// 测试工程
	TestcaseProject *string `json:"testcase_project,omitempty"`

	// 脚本路径
	SvnScriptPath *string `json:"svn_script_path,omitempty"`

	// 约束条件
	MapRestrict *string `json:"map_restrict,omitempty"`

	// 网络脚本名
	NetworkScriptName *string `json:"network_script_name,omitempty"`

	// 自动化类型，非自动化:0, 是自动化:1
	AutoType *int32 `json:"auto_type,omitempty"`

	// 被自动化执行
	ToBeAutoExec *int32 `json:"to_be_auto_exec,omitempty"`

	// 最后一次结果
	LastResult *string `json:"last_result,omitempty"`

	// 最后一次结果Uri
	LastResultUri *string `json:"last_result_uri,omitempty"`

	// 目录Uri
	FeatureUri *string `json:"feature_uri,omitempty"`

	// 测试接口名
	InterfaceName *string `json:"interface_name,omitempty"`

	// 网络问题ID
	SnpNo *string `json:"snp_no,omitempty"`

	// 关联需求编号
	DrRelationId *string `json:"dr_relation_id,omitempty"`

	// 测试基数
	TestBaseNum *string `json:"test_base_num,omitempty"`

	// 是否被自动化执行
	AutomaticallyExecuted *int32 `json:"automatically_executed,omitempty"`

	// 第一次执行时间
	FirstExecuteTime *sdktime.SdkTime `json:"first_execute_time,omitempty"`

	// 检测类型
	DetectType *string `json:"detect_type,omitempty"`

	// 执行参数
	ExecuteParam *string `json:"execute_param,omitempty"`

	// 分析领域
	TestFeature *string `json:"test_feature,omitempty"`

	// 是否是契约用例，0:表示非契约用例, 1：表示契约用例
	IsContractTestcase *int32 `json:"is_contract_testcase,omitempty"`

	// 总共耗时
	TimeCost *float64 `json:"time_cost,omitempty"`

	// 自定义字段1
	CustomField1 *string `json:"custom_field_1,omitempty"`

	// 自定义字段2
	CustomField2 *string `json:"custom_field_2,omitempty"`

	// 自定义字段3
	CustomField3 *string `json:"custom_field_3,omitempty"`

	// 自定义字段4
	CustomField4 *string `json:"custom_field_4,omitempty"`

	// 自定义字段5
	CustomField5 *string `json:"custom_field_5,omitempty"`

	// 自定义字段6
	CustomField6 *string `json:"custom_field_6,omitempty"`

	// 自定义字段7
	CustomField7 *string `json:"custom_field_7,omitempty"`

	// 自定义字段8
	CustomField8 *string `json:"custom_field_8,omitempty"`

	// 自定义字段9
	CustomField9 *string `json:"custom_field_9,omitempty"`

	// 自定义字段10
	CustomField10 *string `json:"custom_field_10,omitempty"`

	// 自定义字段11
	CustomField11 *string `json:"custom_field_11,omitempty"`

	// 自定义字段12
	CustomField12 *string `json:"custom_field_12,omitempty"`

	// 自定义字段13
	CustomField13 *string `json:"custom_field_13,omitempty"`

	// 自定义字段14
	CustomField14 *string `json:"custom_field_14,omitempty"`

	// 自定义字段15
	CustomField15 *string `json:"custom_field_15,omitempty"`

	// 自定义字段16
	CustomField16 *string `json:"custom_field_16,omitempty"`

	// 自定义字段17
	CustomField17 *string `json:"custom_field_17,omitempty"`

	// 自定义字段18
	CustomField18 *string `json:"custom_field_18,omitempty"`

	// 自定义字段19
	CustomField19 *string `json:"custom_field_19,omitempty"`

	// 自定义字段20
	CustomField20 *string `json:"custom_field_20,omitempty"`

	// 自定义字段21
	CustomField21 *string `json:"custom_field_21,omitempty"`

	// 自定义字段22
	CustomField22 *string `json:"custom_field_22,omitempty"`

	// 自定义字段23
	CustomField23 *string `json:"custom_field_23,omitempty"`

	// 自定义字段24
	CustomField24 *string `json:"custom_field_24,omitempty"`

	// 自定义字段25
	CustomField25 *string `json:"custom_field_25,omitempty"`

	// 记录用例由非自动化变为自动化类型的时间
	BeAutoTypeTime *sdktime.SdkTime `json:"be_auto_type_time,omitempty"`

	// 配对用例编号
	CompareNumber *string `json:"compare_number,omitempty"`

	// 场景标识
	SceneFlag *string `json:"scene_flag,omitempty"`

	// 场景标识
	BaseFlag *string `json:"base_flag,omitempty"`

	// 区别是否从yaml中生成的用例，默认false
	ParaValidator *string `json:"para_validator,omitempty"`

	// knet节点id
	KnetNodeId *string `json:"knet_node_id,omitempty"`

	// 最后一次执行用户
	LastExeAuthor *string `json:"last_exe_author,omitempty"`

	// 运营商
	CloudCarrier *string `json:"cloud_carrier,omitempty"`

	// 应用市场
	MarketPlace *string `json:"market_place,omitempty"`

	// 脑图id
	TestMindId *string `json:"test_mind_id,omitempty"`

	// 脑图url
	TestMindUrl *string `json:"test_mind_url,omitempty"`

	// git提交url
	CommitUrl *string `json:"commit_url,omitempty"`

	// 测试模式编号
	TestPatternNumber *string `json:"test_pattern_number,omitempty"`

	// 测试因子编号
	TestFactorNumber *string `json:"test_factor_number,omitempty"`

	// 状态Code
	StatusCode *string `json:"status_code,omitempty"`

	// 结果Code
	ResultCode *string `json:"result_code,omitempty"`

	// 迭代ID
	ReleaseId *string `json:"release_id,omitempty"`

	// 标签ID
	LabelId *string `json:"label_id,omitempty"`

	// 对外用例操作时，标签名列表
	LabelNames *[]string `json:"label_names,omitempty"`

	// 模块ID
	ModuleId *string `json:"module_id,omitempty"`

	// 最后执行时间
	ExecuteLatestTime *sdktime.SdkTime `json:"execute_latest_time,omitempty"`

	// 执行时长
	ExecuteDuration *string `json:"execute_duration,omitempty"`

	// 是否关键用例
	IsKeyword *int32 `json:"is_keyword,omitempty"`

	// 测试版本号
	ReleaseDev *string `json:"release_dev,omitempty"`

	// 是否用户新增用例
	NewCreated *int32 `json:"new_created,omitempty"`

	// 执行参数
	ExecuteParameter *string `json:"execute_parameter,omitempty"`

	// 项目ID，外部使用项目ID，内部使用默认值
	ProjectUuid *string `json:"project_uuid,omitempty"`

	// 分支或者迭代uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 更新用例信息列表
	CaseList *[]CaseInfo `json:"case_list,omitempty"`

	// 批量更新用例id列表
	CaseIdList *[]string `json:"case_id_list,omitempty"`
}

func (o UpdateTestCaseListInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTestCaseListInfo struct{}"
	}

	return strings.Join([]string{"UpdateTestCaseListInfo", string(data)}, " ")
}
