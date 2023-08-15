package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Graph1 struct {

	// 图ID。
	Id *string `json:"id,omitempty"`

	// 图名称。
	Name *string `json:"name,omitempty"`

	// 图的创建人账号。
	CreatedBy *string `json:"createdBy,omitempty"`

	// 是否支持跨AZ高可用。
	IsMultiAz *string `json:"isMultiAz,omitempty"`

	// 域编码。
	RegionCode *string `json:"regionCode,omitempty"`

	// 可用区编码。
	AzCode *string `json:"azCode,omitempty"`

	// 元数据文件路径。
	SchemaPath *[]SchemaPath1 `json:"schemaPath,omitempty"`

	// 边数据集OBS路径。
	EdgesetPath *[]EdgesetPath1 `json:"edgesetPath,omitempty"`

	// 边数据集文件格式。
	EdgesetFormat *string `json:"edgesetFormat,omitempty"`

	// 边数据集文件默认Label。
	EdgesetDefaultLabel *string `json:"edgesetDefaultLabel,omitempty"`

	// 点数据集OBS路径。
	VertexsetPath *[]VertexsetPath1 `json:"vertexsetPath,omitempty"`

	// 点数据集文件格式。
	VertexsetFormat *string `json:"vertexsetFormat,omitempty"`

	// 点数据集文件默认Label。
	VertexsetDefaultLabel *string `json:"vertexsetDefaultLabel,omitempty"`

	// 图版本。
	DataStoreVersion *string `json:"dataStoreVersion,omitempty"`

	// 企业项目信息，如果未指定则不开启，默认不开启。
	SysTags *[]string `json:"sys_tags,omitempty"`

	// 图的状态码。  - 100：准备中 - 200：运行中 - 201：升级中 - 202：导入中 - 203：回滚中 - 204：导出中 - 205：清空中 - 206：扩容准备中 - 207：扩容中 - 208：扩容回退中 - 300：故障 - 303：创建失败 - 400：被删除 - 800：已冻结 - 900：停止 - 901：停止中 - 920：启动中
	Status *string `json:"status,omitempty"`

	// 图创建进度百分比。 >只有图状态码为100时返回该字段。
	ActionProgress *string `json:"actionProgress,omitempty"`

	// 图规模类型索引。  - 0：一万边 - 1：百万边 - 2：千万边 - 3：一亿边 - 4：十亿边 - 5：百亿边 - 6：千亿边
	GraphSizeTypeIndex *string `json:"graphSizeTypeIndex,omitempty"`

	// 虚拟私有云ID。
	VpcId *string `json:"vpcId,omitempty"`

	// 指定虚拟私有云下的子网ID。
	SubnetId *string `json:"subnetId,omitempty"`

	// 安全组ID。
	SecurityGroupId *string `json:"securityGroupId,omitempty"`

	// 副本个数，默认为1。
	Replication *int32 `json:"replication,omitempty"`

	// 图创建时间。
	Created *string `json:"created,omitempty"`

	// 图更新时间。
	Updated *string `json:"updated,omitempty"`

	// 图实例私有网络访问地址，通过该IP用户可以通过私有网络中已部署的弹性云服务器对图实例进行访问。
	PrivateIp *string `json:"privateIp,omitempty"`

	// 图实例公网访问地址，通过该IP用户可以从互联网对图实例进行访问。
	PublicIp *string `json:"publicIp,omitempty"`

	// 图实例CPU架构类型，取值为x86_64和aarch64。
	Arch *string `json:"arch,omitempty"`

	// 是否加密。默认值为“false”，默认不加密。
	Encrypted *bool `json:"encrypted,omitempty"`

	// 用户主密钥ID。
	MasterKeyId *string `json:"masterKeyId,omitempty"`

	// 用户主密钥名称。
	MasterKeyName *string `json:"masterKeyName,omitempty"`

	// 是否启用细粒度权限控制。
	EnableRBAC *bool `json:"enableRBAC,omitempty"`

	// 是否启用全文索引。
	EnableFulltextIndex *bool `json:"enableFulltextIndex,omitempty"`

	// 是否启用HyG，该参数只对千亿规格图生效
	EnableHyG *bool `json:"enableHyG,omitempty"`

	// 图实例私有网络访问物理地址列表。为了防止浮动IP切换造成业务闪断，我们推荐您通过轮询的方式使用物理IP访问图实例。
	TrafficIpList *[]string `json:"trafficIpList,omitempty"`

	// 图实例加密算法，取值为：  - generalCipher：国际算法 - SMcompatible：商密算法（兼容国际）
	CryptAlgorithm *string `json:"cryptAlgorithm,omitempty"`

	// 是否开启安全模式，开启安全模式会对性能有较大影响。
	EnableHttps *bool `json:"enableHttps,omitempty"`

	// 标签列表，每个标签用<key,value>键值对表示。
	Tags *[]interface{} `json:"tags,omitempty"`
}

func (o Graph1) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Graph1 struct{}"
	}

	return strings.Join([]string{"Graph1", string(data)}, " ")
}
