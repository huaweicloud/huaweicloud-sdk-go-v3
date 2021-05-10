package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Detailed cluster parameters.
type V3ClusterSpec struct {
	// 集群类型：  - VirtualMachine：CCE集群  基于Kubernetes来管理一组节点资源，支持虚拟机和裸金属的管理，Kubernetes将自动调度容器运行在可用节点上。在创建容器工作负载前，您需要存在一个可用集群。  - ARM64: 鲲鹏集群  鲲鹏容器集群（ARM指令集）提供了容器在鲲鹏（ARM架构）服务器上的运行能力，提供与X86服务器相同的调度伸缩和快速部署能力。

	Type V3ClusterSpecType `json:"type"`
	// 字段默认值：创建CCE集群[或鲲鹏集群](tag:hws)时，如果是非专属云为 cce.s1.small，专属云则为 cce.dec.s1.small；   集群规格，集群创建完成后规格不可再变更，请按实际业务需求进行选择：   - cce.s1.small: 小规模单控制节点CCE集群（最大50节点）  - cce.s1.medium: 中等规模单控制节点CCE集群（最大200节点）  - cce.s2.small: 小规模多控制节点CCE集群（最大50节点）  - cce.s2.medium: 中等规模多控制节点CCE集群（最大200节点）  - cce.s2.large: 大规模多控制节点CCE集群（最大1000节点）  - cce.s2.xlarge: 超大规模多控制节点CCE集群（最大2000节点）   >    - s1：单控制节点CCE集群。 >    - s2：多控制节点CCE集群。 >    - dec：专属CCE集群规格。如cce.dec.s1.small为小规模单控制节点专属CCE集群（最大50节点）。 >    - 最大节点数：当前集群支持管理的最大节点规模，请根据业务需求选择。 >    - 单控制节点集群：普通集群是单控制节点，控制节点故障后，集群将不可用，但已运行工作负载不受影响。 >    - 多控制节点集群：即高可用集群，当某个控制节点故障时，集群仍然可用。查看集群模式请参见[[如何排查已创建的集群是否为高可用集群？](https://support.huaweicloud.com/cce_faq/cce_faq_00155.html)](tag:hws)[[如何排查已创建的集群是否为高可用集群？](https://support.huaweicloud.com/intl/zh-cn/cce_faq/cce_faq_00155.html)](tag:hws_hk)

	Flavor string `json:"flavor"`
	// 集群版本，与Kubernetes社区基线版本保持一致，建议选择最新版本。  在CCE控制台中支持创建两种最新版本的集群。可登录CCE控制台，单击“总览 > 购买Kubernetes集群”，在“版本”处获取到集群版本。 其它集群版本，当前仍可通过api创建，但后续会逐渐下线，具体下线策略请关注CCE官方公告。  >    - 若不配置，默认创建最新版本的集群。 >    - 若指定集群基线版本但是不指定具体r版本，则系统默认选择对应集群版本的最新r版本。建议不指定具体r版本由系统选择最新版本。

	Version *string `json:"version,omitempty"`
	// 集群描述，对于集群使用目的的描述，可根据实际情况自定义，默认为空。集群创建成功后可通过接口[[更新指定的集群](https://support.huaweicloud.com/api-cce/cce_02_0240.html)](tag:hws)[[更新指定的集群](https://support.huaweicloud.com/intl/zh-cn/api-cce/cce_02_0240.html)](tag:hws_hk)来做出修改，也可在CCE控制台中对应集群的“集群详情”下的“描述”处进行修改。仅支持utf-8编码。

	Description *string `json:"description,omitempty"`
	// 集群是否使用IPv6模式，1.15版本及以上支持。

	Ipv6enable *bool `json:"ipv6enable,omitempty"`
	// CCE Turbo集群(公测)

	OffloadCluster *bool `json:"offloadCluster,omitempty"`

	HostNetwork *HostNetwork `json:"hostNetwork"`

	ContainerNetwork *ContainerNetwork `json:"containerNetwork"`

	EniNetwork *EniNetwork `json:"eniNetwork,omitempty"`

	Authentication *Authentication `json:"authentication,omitempty"`
	// 集群的计费方式，当前接口只支持创建“按需计费”的集群。计费方式为“按需计费”时，取值为“0”。若不填，则默认为“0”。

	BillingMode *int32 `json:"billingMode,omitempty"`
	// 控制节点的高级配置

	Masters *[]MasterSpec `json:"masters,omitempty"`
	// 服务网段参数，kubernetes clusterIp取值范围，1.11.7版本及以上支持。

	KubernetesSvcIpRange *string `json:"kubernetesSvcIpRange,omitempty"`
	// 集群资源标签

	ClusterTags *[]ResourceTag `json:"clusterTags,omitempty"`
	// 服务转发模式，支持以下两种实现：  iptables：社区传统的kube-proxy模式，完全以iptables规则的方式来实现service负载均衡。该方式最主要的问题是在服务多的时候产生太多的iptables规则，非增量式更新会引入一定的时延，大规模情况下有明显的性能问题。  ipvs：主导开发并在社区获得广泛支持的kube-proxy模式，采用增量式更新，吞吐更高，速度更快，并可以保证service更新期间连接保持不断开，适用于大规模场景。   >此参数目前仅在响应中体现，创建集群时请在**extendParam**中配置此参数。

	KubeProxyMode *V3ClusterSpecKubeProxyMode `json:"kubeProxyMode,omitempty"`
	// 可用区（仅查询返回字段）

	Az *string `json:"az,omitempty"`
	// 扩展字段，key/value对格式。可配置多可用区集群、专属CCE集群，以及将集群创建在特定的企业项目下。可配置的key/value对如下： - clusterAZ: 集群控制节点可用区配置。   - multi_az：多可用区，可选。仅使用高可用集群时才可以配置多可用区。   - 专属云计算池可用区：用于指定专属云可用区部署集群控制节点。   如果需配置专属CCE集群，该字段为必选。例如“华北四-可用区一”取值为：cn-north-4a。更多信息请参见[[什么是专属计算集群？](https://support.huaweicloud.com/productdesc-dcc/zh-cn_topic_0016310838.html)](tag:hws)[[什么是专属计算集群？](https://support.huaweicloud.com/intl/zh-cn/productdesc-dcc/zh-cn_topic_0016310838.html)](tag:hws_hk) - dssMasterVolumes: 用于指定控制节点的系统盘和数据盘使用专属分布式存储，未指定或者值为空时，默认使用EVS云硬盘。   如果配置专属CCE集群，该字段为必选，请按照如下格式设置：   ```   <rootVol.dssPoolID>.<rootVol.volType>;<dataVol.dssPoolID>.<dataVol.volType>   ```   字段说明：   rootVol为系统盘；dataVol为数据盘；   dssPoolID为专属分布式存储池ID；   volType为专属分布式存储池的存储类型，如SAS、SSD。   样例：c950ee97-587c-4f24-8a74-3367e3da570f.sas;6edbc2f4-1507-44f8-ac0d-eed1d2608d38.ssd   非专属CCE集群不支持配置该字段。 - enterpriseProjectId:   如果需要将集群创建在特定的企业项目下，请设置{\"enterpriseProjectId\":\"xxx\"}的key/value对。   >   - 需要开通企业项目功能后才可配置企业项目，详情请参见 [[如何进入企业管理页面](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0108763975.html)](tag:hws)[[如何进入企业管理页面](https://support.huaweicloud.com/intl/zh-cn/usermanual-em/zh-cn_topic_0108763975.html)](tag:hws_hk)。   >   - 集群所属的企业项目与集群下所关联的其他云服务资源所属的企业项目必须保持一致。 - kubeProxyMode:   服务转发模式，支持以下两种实现：   - iptables：社区传统的kube-proxy模式，完全以iptables规则的方式来实现service负载均衡。该方式最主要的问题是在服务多的时候产生太多的iptables规则，非增量式更新会引入一定的时延，大规模情况下有明显的性能问题   - ipvs：主导开发并在社区获得广泛支持的kube-proxy模式，采用增量式更新，吞吐更高，速度更快，并可以保证service更新期间连接保持不断开，适用于大规模场景。 - clusterExternalIP: master 弹性公网IP - alpha.cce/fixPoolMask:   容器网络固定IP池掩码位数，仅vpc-router支持。   整数字符传取值范围: 24 ~ 28 - decMasterFlavor: 专属CCE集群指定可控制节点的规格。最大长度255。 - dockerUmaskMode: 集群默认Docker的UmaskMode配置，可取值为secure或normal，不指定时默认为normal。 - kubernetes.io/cpuManagerPolicy:   集群CPU管理策略。取值为none或static，默认为none。   - none：关闭工作负载实例独占CPU核的功能，优点是CPU共享池的可分配核数较多   - static：支持给工作负载实例配置CPU独占，适用于对CPU缓存和调度延迟敏感的工作负载。 - periodType:   订购周期类型，取值范围：     - month：月     - year：年   > billingMode为1（包周期）时生效，且为必选。 - periodNum:   订购周期数，取值范围：     - periodType=month（周期类型为月）时，取值为[1-9]。     - periodType=year（周期类型为年）时，取值为1。   > billingMode为1时生效，且为必选。 - isAutoRenew:   是否自动续订     - “true”：自动续订     - “false”：不自动续订   > billingMode为1时生效，不填写此参数时默认不会自动续费。 - isAutoPay:   是否自动扣款     - “true”：自动扣款     - “false”：不自动扣款   > billingMode为1时生效，不填写此参数时默认不会自动扣款。

	ExtendParam map[string]string `json:"extendParam,omitempty"`
	// 支持Istio

	SupportIstio *bool `json:"supportIstio,omitempty"`
}

func (o V3ClusterSpec) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "V3ClusterSpec struct{}"
	}

	return strings.Join([]string{"V3ClusterSpec", string(data)}, " ")
}

type V3ClusterSpecType struct {
	value string
}

type V3ClusterSpecTypeEnum struct {
	VIRTUAL_MACHINE V3ClusterSpecType
	ARM64           V3ClusterSpecType
}

func GetV3ClusterSpecTypeEnum() V3ClusterSpecTypeEnum {
	return V3ClusterSpecTypeEnum{
		VIRTUAL_MACHINE: V3ClusterSpecType{
			value: "VirtualMachine",
		},
		ARM64: V3ClusterSpecType{
			value: "ARM64",
		},
	}
}

func (c V3ClusterSpecType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *V3ClusterSpecType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type V3ClusterSpecKubeProxyMode struct {
	value string
}

type V3ClusterSpecKubeProxyModeEnum struct {
	IPTABLES V3ClusterSpecKubeProxyMode
	IPVS     V3ClusterSpecKubeProxyMode
}

func GetV3ClusterSpecKubeProxyModeEnum() V3ClusterSpecKubeProxyModeEnum {
	return V3ClusterSpecKubeProxyModeEnum{
		IPTABLES: V3ClusterSpecKubeProxyMode{
			value: "iptables",
		},
		IPVS: V3ClusterSpecKubeProxyMode{
			value: "ipvs",
		},
	}
}

func (c V3ClusterSpecKubeProxyMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *V3ClusterSpecKubeProxyMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
