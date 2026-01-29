package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductEntity struct {

	// **参数解释**： 产品类型。 **约束限制**： 不涉及。 **取值范围**： [- single：4.8.0单机。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- cluster：4.8.0集群。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- single.basic：5.x单机基础版。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) [- cluster.basic：5.x集群基础版。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) [- cluster.professional：5.x集群专业版。](tag:hws,hws_eu,hws_hk,ctc,srg) **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： RocketMQ实例规格。[x86环境后缀为.x86，arm环境为.arm。](tag:hcs,fcs,hcs_oemout)[该接口与各引擎共用，因此还会返回RocketMQ其他版本产品规格。](tag:dt)  **约束限制**： 不涉及。  **取值范围**： [当“type”为“single.basic”选择单机规格；当“type”为“cluster.basic”选择集群规格。](tag:dt) [- c6.2u8g.cluster.x86或c6.2u8g.cluster.arm：单个代理最大分区数50，单个代理最大消费组数100](tag:fcs) [- c6.4u8g.cluster.small：单个代理最大Topic数2000，单个代理最大消费组数2000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- c6.4u8g.cluster：单个代理最大Topic数4000，单个代理最大消费组数4000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- c6.4u16g.cluster.x86或c6.4u16g.cluster.arm：单个代理最大分区数100，单个代理最大消费组数200](tag:hcs,hcs_oemout) [- c6.8u16g.cluster：单个代理最大Topic数8000，单个代理最大消费组数8000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- c6.8u32g.cluster.x86或c6.8u32g.cluster.arm：单个代理最大Topic数200，单个代理最大消费组数400](tag:hcs,fcs,hcs_oemout) [- c6.12u24g.cluster：单个代理最大Topic数12000，单个代理最大消费组数12000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- c6.16u64g.cluster.x86或c6.16u64g.cluster.arm：单个代理最大Topic数300，单个代理最大消费组数600](tag:hcs,fcs,hcs_oemout) [- c6.16u32g.cluster：单个代理最大Topic数16000，单个代理最大消费组数16000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- c6.32u128g.cluster.x86或c6.32u128g.cluster.arm：单个代理最大Topic数400，单个代理最大消费组数800](tag:hcs,fcs,hcs_oemout) [- rocketmq.b1.large.1：RocketMQ 5.x 基础版单机规格，实例TPS 500](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) [- rocketmq.b2.large.4：RocketMQ 5.x 基础版集群规格，实例TPS 2000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) [- rocketmq.b2.large.8：RocketMQ 5.x 基础版集群规格，实例TPS 4000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) [- rocketmq.b2.large.12：RocketMQ 5.x 基础版集群规格，实例TPS 6000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) [- rocketmq.p1.large.1：RocketMQ 5.x 专业版单机规格，实例TPS 500](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.8：RocketMQ 5.x 专业版集群规格，实例TPS 4000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.12：RocketMQ 5.x 专业版集群规格，实例TPS 6000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.20：RocketMQ 5.x 专业版集群规格，实例TPS 10000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.40：RocketMQ 5.x 专业版集群规格，实例TPS 20000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.100：RocketMQ 5.x 专业版集群规格，实例TPS 50000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.150：RocketMQ 5.x 专业版集群规格，实例TPS 75000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.200：RocketMQ 5.x 专业版集群规格，实例TPS 100000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.300：RocketMQ 5.x 专业版集群规格，实例TPS 150000](tag:hws,hws_eu,hws_hk,ctc,srg) **默认取值**： 不涉及。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释**： ECS的Flavor ID。[可参考ECS服务实例规格中的规格清单，DMS服务可适配，取通用计算增强型C6系列及上，RocketMQ 5.x CPU/内存配比为1:4。](tag:ctc,hws_eu,ocb,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,hcs,fcs,dt,hcs_oemout,srg) **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EcsFlavorId *string `json:"ecs_flavor_id,omitempty"`

	// **参数解释**： CBC的规格码。 **约束限制**： 不涉及。 **取值范围**： [- dms.platinum.c6：华为云账单计费类型。](tag:hws,hws_hk)[账单计费类型。](tag:ctc,hws_eu,ocb,g42,hk_g42,tm,sbc,hk_sbc,cmcc,hk_tm,hcs,fcs,hcs_oemout,ax,srg) [- dms.rocketmq.basic.single.tps：RocketMQ 5.x基础版单机计费类型。](tag:ctc,hws_eu,ocb,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,hcs,fcs,dt,hcs_oemout,srg) [- dms.rocketmq.basic.cluster.tps：RocketMQ 5.x基础版集群计费类型。](tag:ctc,hws_eu,ocb,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,hcs,fcs,dt,hcs_oemout,srg) [- dms.rocketmq.pro.single.tps：RocketMQ 5.x专业版单机计费类型。](tag:hws,hws_eu,hws_hk,ctc,srg) [- dms.rocketmq.pro.cluster.tps：RocketMQ 5.x专业版集群计费类型。](tag:hws,hws_eu,hws_hk,ctc,srg) [- dms.platinum.c6.dec：专属云账单计费类型。](tag:hws,hws_hk) **默认取值**： 不涉及。
	BillingCode *string `json:"billing_code,omitempty"`

	// **参数解释**： 支持的架构类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ArchTypes *[]string `json:"arch_types,omitempty"`

	// **参数解释**： 计费模式。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ChargingMode *[]string `json:"charging_mode,omitempty"`

	// **参数解释**： 支持的io类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Ios *[]IosEntity `json:"ios,omitempty"`

	// **参数解释**： 支持的特性功能列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SupportFeatures *[]string `json:"support_features,omitempty"`

	Properties *ProductEntityProperties `json:"properties,omitempty"`

	// **参数解释**： 是否为擎天实例。 **约束限制**： 不涉及。 **取值范围**： - true：是 - false：否 **默认取值**： 不涉及。
	QingtianIncompatible *bool `json:"qingtian_incompatible,omitempty"`
}

func (o ProductEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductEntity struct{}"
	}

	return strings.Join([]string{"ProductEntity", string(data)}, " ")
}
