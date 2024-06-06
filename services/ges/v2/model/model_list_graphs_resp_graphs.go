package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListGraphsRespGraphs struct {

	// 图ID。
	Id *string `json:"id,omitempty"`

	// 图名称。
	Name *string `json:"name,omitempty"`

	// 图的创建人账号。
	CreatedBy *string `json:"created_by,omitempty"`

	// 是否支持跨AZ高可用。
	IsMultiAz *string `json:"is_multi_az,omitempty"`

	// 域编码。
	RegionCode *string `json:"region_code,omitempty"`

	// 可用区编码。
	AzCode *string `json:"az_code,omitempty"`

	// 元数据文件路径。
	SchemaPath *[]ListGraphsRespSchemaPath `json:"schema_path,omitempty"`

	// 边数据集OBS路径。
	EdgesetPath *[]ListGraphsRespSchemaPath `json:"edgeset_path,omitempty"`

	// 点数据集OBS路径。
	VertexsetPath *[]ListGraphsRespSchemaPath `json:"vertexset_path,omitempty"`

	// 边数据集文件格式。
	EdgesetFormat *string `json:"edgeset_format,omitempty"`

	// 边数据集文件默认Label。
	EdgesetDefaultLabel *string `json:"edgeset_default_label,omitempty"`

	// 点数据集文件格式。
	VertexsetFormat *string `json:"vertexset_format,omitempty"`

	// 点数据集文件默认Label。
	VertexsetDefaultLabel *string `json:"vertexset_default_label,omitempty"`

	// 图版本。
	DataStoreVersion *string `json:"data_store_version,omitempty"`

	// 企业项目信息，如果未指定则不开启，默认不开启。
	SysTags *[]string `json:"sys_tags,omitempty"`

	// 图的状态码。  - 100：准备中 - 200：运行中 - 201：升级中 - 202：导入中 - 203：回滚中 - 204：导出中 - 205：清空中 - 206：扩容准备中 - 207：扩容中 - 208：扩容回退中 - 210：扩副本准备中 - 211：扩副本中 - 300：故障 - 303：创建失败 - 400：被删除 - 800：已冻结 - 900：停止 - 901：停止中 - 920：启动中
	Status *string `json:"status,omitempty"`

	// 图创建进度百分比。 > 只有图状态码为100时返回该字段。
	ActionProgress *string `json:"action_progress,omitempty"`

	// 图规模类型索引。  - 0：一万边 - 1：百万边 - 2：千万边 - 3：一亿边 - 4：十亿边 - 5：百亿边 - 6：千亿边 - 401：十亿增强边
	GraphSizeTypeIndex *string `json:"graph_size_type_index,omitempty"`

	// 虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 指定虚拟私有云下的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 副本个数，默认为1。
	Replication *int32 `json:"replication,omitempty"`

	// 图创建时间。
	Created *string `json:"created,omitempty"`

	// 图更新时间。
	Updated *string `json:"updated,omitempty"`

	// 图实例私有网络访问浮动IP地址，通过该IP用户可以通过私有网络中已部署的弹性云服务器对图实例进行访问。
	PrivateIp *string `json:"private_ip,omitempty"`

	// 图实例公网访问地址，通过该IP用户可以从互联网对图实例进行访问。
	PublicIp *string `json:"public_ip,omitempty"`

	// 图实例CPU架构类型，取值为x86_64和aarch64。
	Arch *string `json:"arch,omitempty"`

	// 是否加密。默认值为“false”，默认不加密。
	Encrypted *bool `json:"encrypted,omitempty"`

	// 用户主密钥ID。
	MasterKeyId *string `json:"master_key_id,omitempty"`

	// 用户主密钥名称。
	MasterKeyName *string `json:"master_key_name,omitempty"`

	// 是否启用细粒度权限控制。
	EnableRbac *bool `json:"enable_rbac,omitempty"`

	// 是否启用全文索引。
	EnableFullTextIndex *bool `json:"enable_full_text_index,omitempty"`

	// 是否启用HyG，该参数只对千亿规格图生效。
	EnableHyg *bool `json:"enable_hyg,omitempty"`

	// 图实例私有网络访问物理地址列表。为了防止浮动IP切换造成业务闪断，我们推荐您通过轮询的方式使用物理IP访问图实例。
	TrafficIpList *[]string `json:"traffic_ip_list,omitempty"`

	// 图实例加密算法，取值为：  - generalCipher：国密算法 - SMcompatible：商密算法（兼容国际）
	CryptAlgorithm *string `json:"crypt_algorithm,omitempty"`

	// 是否开启安全模式，开启安全模式会对性能有较大影响
	EnableHttps *bool `json:"enable_https,omitempty"`

	// 标签列表，每个标签用<key,value>键值对表示。
	Tags *[]ListGraphsRespTags `json:"tags,omitempty"`

	// 图产品类型，取值为InMemory和Persistence，默认为InMemory，当graph_size_type_index取值为\"6\"时，默认为Persistence。  - InMemory：内存版 - Persistence：持久化版
	ProductType *string `json:"product_type,omitempty"`

	VertexIdType *ListGraphsRespVertexIdType `json:"vertex_id_type,omitempty"`

	// 图的初始规格。该参数从2.3.15版本后开始支持。
	OriginGraphSizeTypeIndex *string `json:"origin_graph_size_type_index,omitempty"`

	// 图扩副本的时间。
	ExpandTime *string `json:"expand_time,omitempty"`

	// 图扩容的时间。
	ResizeTime *string `json:"resize_time,omitempty"`

	// 是否启用多标签。
	EnableMultiLabel *bool `json:"enable_multi_label,omitempty"`

	// 图的容量倍率。只有持久化版百亿规格图支持该参数，该参数从2.3.18版本后开始支持。
	CapacityRatio *int32 `json:"capacity_ratio,omitempty"`

	// 图的sortKey类型，内存版图无此值。
	SortKeyType *string `json:"sort_key_type,omitempty"`
}

func (o ListGraphsRespGraphs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphsRespGraphs struct{}"
	}

	return strings.Join([]string{"ListGraphsRespGraphs", string(data)}, " ")
}
