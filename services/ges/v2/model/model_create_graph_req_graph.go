package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 图类型。
type CreateGraphReqGraph struct {

	// 图名称（输入长度在4位到50位之间，必须以字母开头，可以包含字母、数字或者下划线，不能包含其他的特殊字符）。
	Name string `json:"name"`

	// 图规模类型索引。  - 0：一万边 - 1：百万边 - 2：千万边 - 3：一亿边 - 4：十亿边 - 5：百亿边 - 6：千亿边 - 401：十亿增强边
	GraphSizeTypeIndex string `json:"graph_size_type_index"`

	// 图实例CPU架构类型，取值为x86_64和aarch64。默认取x86_64。  - x86_64：X64 64位架构。 - aarch64：ARM 64位架构。
	Arch *string `json:"arch,omitempty"`

	// 虚拟私有云ID。
	VpcId string `json:"vpc_id"`

	// 指定虚拟私有云下的子网ID。
	SubnetId string `json:"subnet_id"`

	// 安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	PublicIp *CreateGraphReqGraphPublicIp `json:"public_ip,omitempty"`

	// 创建的图是否支持跨可用区（AZ），默认值是false，如果设置为true，系统将会把图中的ECS建在两个可用区中。  如果创建图时，不加该参数，则会将图中的ECS都建在一个可用区中。
	EnableMultiAz *bool `json:"enable_multi_az,omitempty"`

	Encryption *CreateGraphReqGraphEncryption `json:"encryption,omitempty"`

	LtsOperationTrace *CreateGraphReqGraphLtsOperationTrace `json:"lts_operation_trace,omitempty"`

	// 企业项目信息，如果未指定则不开启，默认不开启。
	SysTags *[]CreateGraphReqGraphSysTags `json:"sys_tags,omitempty"`

	// 支持标签TMS，做费用归集，默认不开启。
	Tags *[]ListGraphsRespTags `json:"tags,omitempty"`

	// 创建的图是否启用细粒度权限控制，默认不启用，值为false。如果设置为true，创建的图所有用户都没有权限，需要调用业务面细粒度权限控制API进行授权操作才可以访问图。
	EnableRbac *bool `json:"enable_rbac,omitempty"`

	// 创建的图是否开启全文索引控制，默认不启用，值为false。 如果设置为true，十亿增强版-规格版图支持全文索引，创建图时会创建云搜索服务集群。 > 开启全文索引功能。如果CSS服务已经部署，图实例会自动创建CSS集群，图创建时间较长。如果CSS服务没有部署则图创建失败。
	EnableFullTextIndex *bool `json:"enable_full_text_index,omitempty"`

	// 该参数只对千亿规格图生效。
	EnableHyg *bool `json:"enable_hyg,omitempty"`

	// 图实例加密算法，取值为：  - generalCipher：国密算法 - SMcompatible：商密算法（兼容国际）
	CryptAlgorithm string `json:"crypt_algorithm"`

	// 是否开启安全模式，开启安全模式会对性能有较大影响
	EnableHttps bool `json:"enable_https"`

	// 图产品类型，取值为InMemory和Persistence，默认为InMemory，当graph_size_type_index取值为\"6\"时，默认为Persistence。  - InMemory：内存版 - Persistence：持久化版
	ProductType *string `json:"product_type,omitempty"`

	VertexIdType *CreateGraphReqGraphVertexIdType `json:"vertex_id_type,omitempty"`
}

func (o CreateGraphReqGraph) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraphReqGraph struct{}"
	}

	return strings.Join([]string{"CreateGraphReqGraph", string(data)}, " ")
}
