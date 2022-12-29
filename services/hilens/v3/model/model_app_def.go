package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppDef struct {

	// app应用的订单ID，技能来源是市场时，如果不填，则自动选择默认订单。
	AppOrderId *string `json:"app_order_id,omitempty"`

	// app应用的地址，可以是镜像地址或者OBS地址
	AppUrl string `json:"app_url"`

	// 路数限制，添加作业的时候，摄像头和VCN的最大路数不超过该值，范围是0到1000
	ChannelLimit *int32 `json:"channel_limit,omitempty"`

	// 用户可以指定的路数限制上限，范围是0到1000
	ChannelUpperLimit *int32 `json:"channel_upper_limit,omitempty"`

	// 容器启动参数，字符总长度最大为65536，不允许^#~^$|%&*<>()'\"\\[\\]{}]等特殊字符
	Args *[]string `json:"args,omitempty"`

	// 容器启动命令，字符总长度最大为65536。 command支持使用数组定义多条命令，但在IEF控制台界面只会显示第一条命令。不允许^#~^$|%&*<>()'\"\\[\\]{}]等特殊字符
	Command *[]string `json:"command,omitempty"`

	// 环境变量
	Envs *[]Env `json:"envs,omitempty"`

	// 是否是modelbox镜像
	IsModelbox *bool `json:"is_modelbox,omitempty"`

	LivenessProbe *Probe `json:"liveness_probe,omitempty"`

	// 消息变量
	Msgs *[]Env `json:"msgs,omitempty"`

	// 应用名字，只允许英文小写字母、数字、中划线，最大长度32, 英文小写字母或数字开头和结尾。该名称同时对应技能名称，当不传订单id的时候，默认通过该名称和版本号version字段，选择指定技能版本，进行部署，并选择可用的订单（默认订单优先）扣除份额。
	Name string `json:"name"`

	// npu类型，支持D310类型和D910类型。D310表示D310类型。  D910表示D910类型。不填表示为D310类型
	NpuType *string `json:"npu_type,omitempty"`

	// 容器端口映射值
	Ports *[]HostContainerPortMapping `json:"ports,omitempty"`

	// 是否启用特权容器,默认值false
	Privileged *bool `json:"privileged,omitempty"`

	ReadinessProbe *Probe `json:"readiness_probe,omitempty"`

	Resources *ResQuest `json:"resources"`

	// 版本号，长度不操作128，支持大小写数字，下划线，点，中划线
	Version *string `json:"version,omitempty"`

	// 卷配置
	Volumes *[]Volume `json:"volumes,omitempty"`

	StartResources *ResQuest `json:"start_resources,omitempty"`

	ChannelResources *ResQuest `json:"channel_resources,omitempty"`

	// 技能管理ID，技能来源source是skill的时候，需要传入该ID
	SkillProjectId *string `json:"skill_project_id,omitempty"`
}

func (o AppDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppDef struct{}"
	}

	return strings.Join([]string{"AppDef", string(data)}, " ")
}
