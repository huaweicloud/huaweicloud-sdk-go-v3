package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmssTestcase struct {

	// 活动id
	ActivityId *string `json:"activity_id,omitempty"`

	// 应用产品
	AppliedProduct *string `json:"applied_product,omitempty"`

	// 创建人
	Author *string `json:"author,omitempty"`

	// 自动化类型
	AutoType *string `json:"auto_type,omitempty"`

	// 用例分类ID
	CataId *string `json:"cataId,omitempty"`

	// 创建日期
	CreationDate *string `json:"creation_date,omitempty"`

	CustomField1 *string `json:"custom_field_1,omitempty"`

	CustomField10 *string `json:"custom_field_10,omitempty"`

	CustomField11 *string `json:"custom_field_11,omitempty"`

	CustomField12 *string `json:"custom_field_12,omitempty"`

	CustomField13 *string `json:"custom_field_13,omitempty"`

	CustomField14 *string `json:"custom_field_14,omitempty"`

	CustomField15 *string `json:"custom_field_15,omitempty"`

	CustomField16 *string `json:"custom_field_16,omitempty"`

	CustomField17 *string `json:"custom_field_17,omitempty"`

	CustomField18 *string `json:"custom_field_18,omitempty"`

	CustomField19 *string `json:"custom_field_19,omitempty"`

	CustomField2 *string `json:"custom_field_2,omitempty"`

	CustomField20 *string `json:"custom_field_20,omitempty"`

	CustomField21 *string `json:"custom_field_21,omitempty"`

	CustomField22 *string `json:"custom_field_22,omitempty"`

	CustomField23 *string `json:"custom_field_23,omitempty"`

	CustomField24 *string `json:"custom_field_24,omitempty"`

	CustomField25 *string `json:"custom_field_25,omitempty"`

	CustomField3 *string `json:"custom_field_3,omitempty"`

	CustomField4 *string `json:"custom_field_4,omitempty"`

	CustomField5 *string `json:"custom_field_5,omitempty"`

	CustomField6 *string `json:"custom_field_6,omitempty"`

	CustomField7 *string `json:"custom_field_7,omitempty"`

	CustomField8 *string `json:"custom_field_8,omitempty"`

	CustomField9 *string `json:"custom_field_9,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// dr关系ID
	DrRelationid *string `json:"dr_relationid,omitempty"`

	// 环境类型
	EnvType *string `json:"env_type,omitempty"`

	// 执行平台
	ExePlatform *string `json:"exe_platform,omitempty"`

	// 内部预期输出
	ExpectOutput *string `json:"expect_output,omitempty"`

	// 特性路径
	FeaturePath *string `json:"feature_path,omitempty"`

	// 是否为关键字
	IsKeyWord *int32 `json:"isKeyWord,omitempty"`

	// 是否为合同测试用例
	IsContractTestcase *string `json:"is_contract_testcase,omitempty"`

	// 是否为参数校验测试用例
	IsParaValidatorTestcase *string `json:"is_paraValidator_testcase,omitempty"`

	// 标签ID列表
	LabelId *[]string `json:"labelId,omitempty"`

	// 最后修改时间
	LastModified *string `json:"last_modified,omitempty"`

	// 最后修改人
	LastModifier *string `json:"last_modifier,omitempty"`

	// 最后的结果
	LastResult *string `json:"last_result,omitempty"`

	// 用例级别
	Level *int32 `json:"level,omitempty"`

	// 市场
	Market *string `json:"market,omitempty"`

	// 模块ID
	ModuleId *string `json:"moduleId,omitempty"`

	// 用例名称
	Name *string `json:"name,omitempty"`

	// 公共aw和项目的关联关系
	NetworkScriptName *string `json:"networkScriptName,omitempty"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty"`

	// 用例编号
	Number *string `json:"number,omitempty"`

	// 原始的uri
	OriginUri *string `json:"originUri,omitempty"`

	// 测试的所有者
	Owner *string `json:"owner,omitempty"`

	// 所有者ID
	OwnerId *string `json:"ownerId,omitempty"`

	// 测试前置条件
	Preparation *string `json:"preparation,omitempty"`

	// 新服务新建用例版本号
	ReleaseDev *string `json:"releaseDev,omitempty"`

	// 发布ID
	ReleaseId *string `json:"releaseId,omitempty"`

	// 备注
	Remark *string `json:"remark,omitempty"`

	// 阶段
	Stage *string `json:"stage,omitempty"`

	// 存储测试步骤和预期结果值对象
	Steps *[]TmssStep `json:"steps,omitempty"`

	// 脚本路径
	SvnScriptPath *string `json:"svnScriptPath,omitempty"`

	// 标签
	Tags *string `json:"tags,omitempty"`

	// 测试特性
	TestFeature *string `json:"test_feature,omitempty"`

	// 内部测试步骤
	TestStep *string `json:"test_step,omitempty"`

	// 测试类型
	TestType *int32 `json:"test_type,omitempty"`

	// 用例基线Uri(BAM 接口覆盖率迭代下用例信息获取用)
	Uri *string `json:"uri,omitempty"`
}

func (o TmssTestcase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmssTestcase struct{}"
	}

	return strings.Join([]string{"TmssTestcase", string(data)}, " ")
}
