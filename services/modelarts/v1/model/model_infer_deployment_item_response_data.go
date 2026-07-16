package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InferDeploymentItemResponseData struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **取值范围：** 服务ID。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 部署名，用户在[创建服务](CreateInferService.xml)时自定义。 **取值范围：** 支持1-128个字符，可以包含字母、汉字、数字、连字符和下划线。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 部署ID。 **取值范围：** 不涉及
	InferName *string `json:"infer_name,omitempty"`

	// **参数解释：** 服务实例组id。 **取值范围：** 不涉及
	ServiceGroupName *string `json:"service_group_name,omitempty"`

	// **参数解释：** 服务当前状态，一次只支持一种状态筛选。默认不过滤。 **取值范围：** - DEPLOYING：部署中。 - FAILED：失败。 - STOPPED：停止。 - RUNNING：运行中。 - DELETING：删除中。 - STOPPING：停止中。 - CONCERNING：告警。 - DELETED：已删除。 - RESTARTING：重启中。 - UPGRADING：升级中。 - ERROR：异常。 - INTERRUPTING：中断中。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 部署对接lts状态。 **取值范围：** - ON：开启。 - OFF：关闭。
	LtsState *string `json:"lts_state,omitempty"`

	// **参数解释：** 是否开启镜像流量。 **取值范围：** 不涉及
	MirrorTrafficEnable *bool `json:"mirror_traffic_enable,omitempty"`

	// **参数解释：** 镜像流量权重。 **取值范围：** 50。
	MirrorTrafficWeight *string `json:"mirror_traffic_weight,omitempty"`

	// **参数解释：** 权重百分比，分配到此模型的流量权重，仅当模型部署为在线服务时需要配置。 **取值范围：** [0, 100]。
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释：** 流量比例，单个部署实例预期接收用户的流量与总流量比值，是由流量权重配置和部署状态计算所得的值。 **取值范围：** 0.00%~100.00%。
	TrafficRatio *string `json:"traffic_ratio,omitempty"`

	// **参数解释：** 专属资源池ID。 **取值范围：** 50。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 服务版本号，以数字及点号组成，形如1.0.1。 **取值范围：** 版本长度不超过8位。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 部署类型。 **取值范围：** - SINGLE：单机单卡。 - MULTI：多机多卡。 - DIST：分布式部署。
	DeployType *string `json:"deploy_type,omitempty"`

	// **参数解释：** 创建时间，根据创建时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释：** 更新时间，根据更新时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释：** 当服务或者部署被冻结时返回的冻结类型信息。
	FrozenInfos *[]FrozenInfo `json:"frozen_infos,omitempty"`
}

func (o InferDeploymentItemResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InferDeploymentItemResponseData struct{}"
	}

	return strings.Join([]string{"InferDeploymentItemResponseData", string(data)}, " ")
}
