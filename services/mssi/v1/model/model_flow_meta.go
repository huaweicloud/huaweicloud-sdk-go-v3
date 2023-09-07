package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlowMeta 编排的Flow元数据
type FlowMeta struct {

	// 发布到apic的api id
	ApicId *string `json:"apic_id,omitempty"`

	// 发布到apic的状态
	ApicReleaseStatus *string `json:"apic_release_status,omitempty"`

	// api流注册到apig的url
	ApigUrl *string `json:"apig_url,omitempty"`

	// 连接器
	Connectors *[]interface{} `json:"connectors,omitempty"`

	// 连接器最新版本
	ConnectorsLatest *[]interface{} `json:"connectors_latest,omitempty"`

	// 创建时间
	CreatdTime *sdktime.SdkTime `json:"creatd_time,omitempty"`

	// 流的描述信息
	Description *string `json:"description,omitempty"`

	// 开发状态
	DevStatus *string `json:"dev_status,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 企业项目名称
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 流/流模板扩展配置列表
	ExtendConfigs *interface{} `json:"extend_configs,omitempty"`

	// 流来源： inner: 公共流模板； custom： 我的流模板
	FlowSourceType *string `json:"flow_source_type,omitempty"`

	// 流/流模板函数列表
	Functions *interface{} `json:"functions,omitempty"`

	// 流的编排数据(大josn)
	Graph *interface{} `json:"graph,omitempty"`

	HisTransferStatus *string `json:"his_transfer_status,omitempty"`

	// logo base64编码
	Icon *string `json:"icon,omitempty"`

	// ID
	Id *string `json:"id,omitempty"`

	// 已部署的connector id
	InstalledConnector *string `json:"installed_connector,omitempty"`

	IsValid *bool `json:"is_valid,omitempty"`

	Label *string `json:"label,omitempty"`

	// 流的名称
	Name *string `json:"name,omitempty"`

	NoticeStatus *string `json:"notice_status,omitempty"`

	// 用户项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// schema版本
	SchemaVersion *string `json:"schemaVersion,omitempty"`

	// 流的状态
	Status *string `json:"status,omitempty"`

	// 流的编排数据（大josn）
	Steps *[]map[string]interface{} `json:"steps,omitempty"`

	// 标签列表
	Tags *[]Tag `json:"tags,omitempty"`

	// 模板ID
	TemplateId *string `json:"template_id,omitempty"`

	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`

	// 测试结果
	TestResult *string `json:"test_result,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 修改时间
	UpdatedTime *sdktime.SdkTime `json:"updated_time,omitempty"`

	// 账号（所有者）
	UserId *string `json:"user_id,omitempty"`

	// 流的版本号
	Version *string `json:"version,omitempty"`

	// webhook触发url的ID
	Webhook *string `json:"webhook,omitempty"`
}

func (o FlowMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowMeta struct{}"
	}

	return strings.Join([]string{"FlowMeta", string(data)}, " ")
}
