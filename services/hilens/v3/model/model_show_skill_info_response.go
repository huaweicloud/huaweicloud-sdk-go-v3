package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSkillInfoResponse Response Object
type ShowSkillInfoResponse struct {

	// 技能应用场景
	SubScenes *string `json:"sub_scenes,omitempty"`

	// 应用模板ID
	AppTemplateId *string `json:"app_template_id,omitempty"`

	// 技能图标
	Icon *string `json:"icon,omitempty"`

	// 技能版本数量
	VersionNum *int32 `json:"version_num,omitempty"`

	// 技能描述
	Description *string `json:"description,omitempty"`

	// 计费编码信息
	ProductInfo *[]string `json:"product_info,omitempty"`

	// 技能类别，分为standard和lite
	Type *string `json:"type,omitempty"`

	// 技能操作系统平台，其值为：Linux，Android， iOS， LiteOS，Windows
	Platform *string `json:"platform,omitempty"`

	// 自研标识，1表示是HiLens自研算法。
	SelfDevFlag *int32 `json:"self_dev_flag,omitempty"`

	// 计费类型，physical_src表示 src
	MeasureType *string `json:"measure_type,omitempty"`

	// 技能审核结果
	ApprovalResult *string `json:"approval_result,omitempty"`

	// 更新时间，形如2022-06-30 17:22:48 GMT+08:00
	UpdateTime *string `json:"update_time,omitempty"`

	// 通道数
	ChannelLimit *int32 `json:"channel_limit,omitempty"`

	// 发布时间
	PublishTime *string `json:"publish_time,omitempty"`

	// 步长
	ResourceStepSize *int32 `json:"resource_step_size,omitempty"`

	// 审批时间
	ApprovalTime *string `json:"approval_time,omitempty"`

	// 云服务编码
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 摘要
	Summary *string `json:"summary,omitempty"`

	// 测试状态
	TestStatus *int32 `json:"test_status,omitempty"`

	// 芯片
	Chip *string `json:"chip,omitempty"`

	// 是否校验模型
	IsVerifyModel *bool `json:"is_verify_model,omitempty"`

	// 技能类型，文件类型file，镜像类型iamge
	Format *string `json:"format,omitempty"`

	// 资源类别
	ResourceType *string `json:"resource_type,omitempty"`

	// 技能版本
	Version *string `json:"version,omitempty"`

	// 计费单位 qps 表示按qps收费，road表示技能路数instance 表示按实例收费
	MeasureUnit *string `json:"measure_unit,omitempty"`

	// 标签
	Tags *[]string `json:"tags,omitempty"`

	// 技能大小
	Size *int32 `json:"size,omitempty"`

	// 测试结果
	TestResult *string `json:"test_result,omitempty"`

	// 安装次数
	InstallTimes *int32 `json:"install_times,omitempty"`

	// 隐私条款
	PrivacyPolicy *[]string `json:"privacy_policy,omitempty"`

	// 技能名字
	Name *string `json:"name,omitempty"`

	// 技能场景
	Scenes *[]string `json:"scenes,omitempty"`

	// 计费模式
	ChargeModel *int32 `json:"charge_model,omitempty"`

	// 云服务资源编码
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 技能Id
	SkillId *string `json:"skill_id,omitempty"`

	// 开发者名字
	Developer *string `json:"developer,omitempty"`

	// 主场景
	MainScenes *string `json:"main_scenes,omitempty"`

	// 所支持的设备类别
	DeviceTypes *[]string `json:"device_types,omitempty"`

	// 技能状态
	Status *int32 `json:"status,omitempty"`

	// 技能版本号列表
	Versions *[]string `json:"versions,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSkillInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSkillInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowSkillInfoResponse", string(data)}, " ")
}
