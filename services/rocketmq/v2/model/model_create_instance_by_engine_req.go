package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceByEngineReq 创建实例请求体。
type CreateInstanceByEngineReq struct {

	// **参数解释**： 实例名称。 **约束限制**： 由英文字符开头，只能由英文字母、数字、中划线、下划线组成，长度为4~64的字符。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 实例的描述信息。 **约束限制**： 长度不超过1024的字符串。[且字符串不能包含\">\"与\"<\"，字符串首字符不能为\"=\",\"+\",\"-\",\"@\"的全角和半角字符。](tag:hcs)  \\与\"在json报文中属于特殊字符，如果参数值中需要显示\\或者\"字符，请在字符前增加转义字符\\，比如\\\\\\\\或者\\\\\"。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 消息引擎类型。 **约束限制**： 不涉及。 **取值范围**： - rocketmq：RocketMQ消息引擎。 - reliability：RocketMQ消息引擎别称。 **默认取值**： 不涉及。
	Engine string `json:"engine"`

	// **参数解释**： 消息引擎的版本。 **约束限制**： 不涉及。 **取值范围**： [- 4.8.0](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- 5.x](tag:hws,hws_eu,hws_hk,ctc,tm,sbc,hk_sbc,hk_tm,hcs,fcs,dt,hcs_oemout,srg) **默认取值**： 不涉及。
	EngineVersion string `json:"engine_version"`

	// **参数解释**： 存储空间，单位：GB。 **取值范围**： 不同的实例规格支持不同的存储配置。 [- RocketMQ 5.x 单机存储取值范围：100-30000。](tag:dt) [- RocketMQ 5.x 集群存储取值范围：200-60000。](tag:dt)
	StorageSpace int32 `json:"storage_space"`

	// **参数解释**： 虚拟私有云ID。获取方法如下：参考《虚拟私有云 API参考》，调用“查询VPC列表”接口，从响应体中获取VPC ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 子网信息。获取方法如下：参考《虚拟私有云 API参考》，调用“查询子网列表”接口，从响应体中获取子网ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SubnetId string `json:"subnet_id"`

	// **参数解释**： 指定实例所属的安全组。获取方法如下：参考《虚拟私有云 API参考》，调用“查询安全组列表”接口，从响应体中获取安全组ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecurityGroupId string `json:"security_group_id"`

	// **参数解释**： 创建节点到指定且有资源的可用区ID。请参考[查询可用区信息](ListAvailableZones.xml)获取可用区ID。 **约束限制**： 该参数不能为空数组或者数组的值为空， 请注意查看该可用区是否有资源。  创建RocketMQ实例，支持节点部署在1个或3个及3个以上的可用区。在为节点指定可用区时，用逗号分隔开。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AvailableZones []string `json:"available_zones"`

	// **参数解释**： RocketMQ实例规格。[x86环境后缀为.x86，arm环境为.arm。single表示单机，cluster表示集群。](tag:hcs,fcs,hcs_oemout)  **约束限制**： 不涉及。  **取值范围**：  [当“type”为“single.basic”选择单机规格；当“type”为“cluster.basic”选择集群规格。](tag:dt) [- c6.2u8g.cluster.x86或c6.2u8g.cluster.arm：单个代理最大分区数50，单个代理最大消费组数100](tag:fcs) [- c6.4u8g.cluster.small：单个代理最大Topic数2000，单个代理最大消费组数2000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- c6.4u8g.cluster：单个代理最大Topic数4000，单个代理最大消费组数4000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- c6.4u16g.cluster.x86或c6.4u16g.cluster.arm：单个代理最大分区数100，单个代理最大消费组数200](tag:hcs,hcs_oemout) [- c6.8u16g.cluster：单个代理最大Topic数8000，单个代理最大消费组数8000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- c6.8u32g.cluster.x86或c6.8u32g.cluster.arm：单个代理最大Topic数200，单个代理最大消费组数400](tag:hcs,fcs,hcs_oemout) [- c6.12u24g.cluster：单个代理最大Topic数12000，单个代理最大消费组数12000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- c6.16u64g.cluster.x86或c6.16u64g.cluster.arm：单个代理最大Topic数300，单个代理最大消费组数600](tag:hcs,fcs,hcs_oemout) [- c6.16u32g.cluster：单个代理最大Topic数16000，单个代理最大消费组数16000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,cmcc,ax,srg) [- c6.32u128g.cluster.x86或c6.32u128g.cluster.arm：单个代理最大Topic数400，单个代理最大消费组数800](tag:hcs,fcs,hcs_oemout) [- rocketmq.b1.large.1：RocketMQ 5.x 基础版单机规格，实例TPS 500](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) [- rocketmq.b2.large.4：RocketMQ 5.x 基础版集群规格，实例TPS 2000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) [- rocketmq.b2.large.8：RocketMQ 5.x 基础版集群规格，实例TPS 4000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) [- rocketmq.b2.large.12：RocketMQ 5.x 基础版集群规格，实例TPS 6000](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,sbc,hk_sbc,hk_tm,dt,srg) [- rocketmq.p1.large.1：RocketMQ 5.x 专业版单机规格，实例TPS 500](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.8：RocketMQ 5.x 专业版集群规格，实例TPS 4000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.12：RocketMQ 5.x 专业版集群规格，实例TPS 6000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.20：RocketMQ 5.x 专业版集群规格，实例TPS 10000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.40：RocketMQ 5.x 专业版集群规格，实例TPS 20000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.100：RocketMQ 5.x 专业版集群规格，实例TPS 50000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.150：RocketMQ 5.x 专业版集群规格，实例TPS 75000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.200：RocketMQ 5.x 专业版集群规格，实例TPS 100000](tag:hws,hws_eu,hws_hk,ctc,srg) [- rocketmq.p2.large.300：RocketMQ 5.x 专业版集群规格，实例TPS 150000](tag:hws,hws_eu,hws_hk,ctc,srg)  **默认取值**： 不涉及。
	ProductId string `json:"product_id"`

	// **参数解释**： 是否打开SSL加密访问。 **约束限制**： 不涉及。 **取值范围**： - true：打开SSL加密访问。 - false：不打开SSL加密访问。 **默认取值**： false。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// **参数解释**： 存储IO规格。 **约束限制**： 不涉及。 **取值范围**： - dms.physical.storage.high.v2：高IO类型磁盘 - dms.physical.storage.ultra.v2：超高IO类型磁盘 [- dms.physical.storage.general：使用通用型SSD的磁盘类型。](tag:hws,hws_hk,dt,ctc,ax) [- dms.physical.storage.extreme：使用极速型SSD的磁盘类型。](tag:hws,hws_hk,dt,ctc,ax) **默认取值**： 不涉及。
	StorageSpecCode string `json:"storage_spec_code"`

	// **参数解释**： 企业项目ID。若为企业项目账号，该参数必填。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 是否开启访问控制列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// **参数解释**： 是否支持IPv6。[华为云Stack不支持此参数。](tag:hcs,hcs_oemout) **约束限制**： 不涉及。 **取值范围**： - true：支持 - false：不支持 **默认取值**： false。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// **参数解释**： 是否开启Proxy功能。 **约束限制**： 不涉及。 **取值范围**： - true：支持 - false：不支持 **默认取值**： false。
	ProxyEnable *bool `json:"proxy_enable,omitempty"`

	// **参数解释**： 是否开启公网访问功能。默认不开启公网。 **约束限制**： 不涉及。 **取值范围**： - true：开启 - false：不开启 **默认取值**： false。
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// **参数解释**： 是否开启磁盘加密。 **约束限制**： 不涉及。 **取值范围**： - true：开启。 - false：不开启。 **默认取值**： false。
	DiskEncryptedEnable *bool `json:"disk_encrypted_enable,omitempty"`

	// **参数解释**： 磁盘加密key，未开启磁盘加密时为空。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DiskEncryptedKey *string `json:"disk_encrypted_key,omitempty"`

	// **参数解释**： 实例绑定的弹性IP地址的ID。 **约束限制**： 以英文逗号隔开多个弹性IP地址的ID。  如果开启了公网访问功能（即enable_publicip为true），该字段为必选。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PublicipId *string `json:"publicip_id,omitempty"`

	// **参数解释**： 代理个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BrokerNum int32 `json:"broker_num"`

	// **参数解释**： 架构类型。 **约束限制**： 不涉及。 **取值范围**： - X86：复杂指令集计算。 - ARM：精简指令集计算。 **默认取值**： 不涉及。
	ArchType *string `json:"arch_type,omitempty"`

	// **参数解释**： 实例使用的安全协议。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TlsMode *string `json:"tls_mode,omitempty"`

	BssParam *BssParam `json:"bss_param,omitempty"`
}

func (o CreateInstanceByEngineReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceByEngineReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceByEngineReq", string(data)}, " ")
}
