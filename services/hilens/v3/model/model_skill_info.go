package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SkillInfo 技能详情
type SkillInfo struct {

	// 技能应用场景
	SubScenes string `json:"sub_scenes"`

	// 应用模板ID
	AppTemplateId string `json:"app_template_id"`

	// 技能图标
	Icon string `json:"icon"`

	// 技能版本数量
	VersionNum int32 `json:"version_num"`

	// 技能描述
	Description string `json:"description"`

	// 计费编码信息
	ProductInfo []string `json:"product_info"`

	// 技能类别，分为standard和lite
	Type string `json:"type"`

	// 技能操作系统平台，其值为：Linux，Android， iOS， LiteOS，Windows
	Platform string `json:"platform"`

	// 自研标识，1表示是HiLens自研算法。
	SelfDevFlag int32 `json:"self_dev_flag"`

	// 计费类型，physical_src表示物理量标，如按周期收费。 src表示一次性
	MeasureType string `json:"measure_type"`

	// 技能审核结果
	ApprovalResult string `json:"approval_result"`

	// 更新时间，形如2022-06-30 17:22:48 GMT+08:00
	UpdateTime string `json:"update_time"`

	// 通道数
	ChannelLimit int32 `json:"channel_limit"`

	// 发布时间
	PublishTime string `json:"publish_time"`

	// 步长
	ResourceStepSize int32 `json:"resource_step_size"`

	// 审批时间
	ApprovalTime string `json:"approval_time"`

	// 云服务编码
	CloudServiceType string `json:"cloud_service_type"`

	// 摘要
	Summary string `json:"summary"`

	// 测试状态
	TestStatus int32 `json:"test_status"`

	// 芯片
	Chip string `json:"chip"`

	// 是否校验模型
	IsVerifyModel bool `json:"is_verify_model"`

	// 技能类型，文件类型file，镜像类型iamge
	Format string `json:"format"`

	// 资源类别
	ResourceType string `json:"resource_type"`

	// 技能版本
	Version string `json:"version"`

	// 计费单位 qps 表示按qps收费，road表示技能路数instance 表示按实例收费
	MeasureUnit string `json:"measure_unit"`

	// 标签
	Tags []string `json:"tags"`

	// 技能大小
	Size int32 `json:"size"`

	// 测试结果
	TestResult string `json:"test_result"`

	// 安装次数
	InstallTimes int32 `json:"install_times"`

	// 隐私条款
	PrivacyPolicy []string `json:"privacy_policy"`

	// 技能名字
	Name string `json:"name"`

	// 技能场景
	Scenes []string `json:"scenes"`

	// 计费模式
	ChargeModel int32 `json:"charge_model"`

	// 云服务资源编码
	ResourceSpecCode string `json:"resource_spec_code"`

	// 技能ID
	SkillId string `json:"skill_id"`

	// 开发者名字
	Developer string `json:"developer"`

	// 主场景
	MainScenes string `json:"main_scenes"`

	// 所支持的设备类别
	DeviceTypes []string `json:"device_types"`

	// 技能状态
	Status int32 `json:"status"`
}

func (o SkillInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SkillInfo struct{}"
	}

	return strings.Join([]string{"SkillInfo", string(data)}, " ")
}
