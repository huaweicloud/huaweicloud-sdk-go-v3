package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Graph
type Graph struct {

	// 图名称（输入长度在4位到50位之间，必须以字母开头，可以包含字母、数字或者下划线，不能包含其他的特殊字符）。
	Name string `json:"name"`

	// 图规模类型索引。 - 0：一万边 - 1：百万边 - 2：千万边 - 3：一亿边 - 4：十亿边 - 5：百亿边 - 401：十亿增强边
	GraphSizeTypeIndex string `json:"graphSizeTypeIndex"`

	// 图实例CPU架构类型，取值为x86_64和aarch64。默认取x86_64。 - x86_64：X64 64位架构。 - aarch64：ARM 64位架构。
	Arch *string `json:"arch,omitempty"`

	DataSource *DataSource `json:"dataSource,omitempty"`

	//   虚拟私有云ID。
	VpcId string `json:"vpcId"`

	// 指定虚拟私有云下的子网ID。
	SubnetId string `json:"subnetId"`

	// 安全组ID。
	SecurityGroupId string `json:"securityGroupId"`

	PublicIp *PublicIp `json:"publicIp,omitempty"`

	// 创建的图是否支持跨可用区（AZ），默认值是false，如果设置为true，系统将会把图中的ECS建在两个可用区中。  如果创建图时，不加该参数，则会将图中的ECS都建在一个可用区中。
	EnableMultiAz *bool `json:"enableMultiAz,omitempty"`

	Encryption *EncryptionReq `json:"encryption,omitempty"`

	LtsOperationTrace *LtsOperationTraceReq `json:"ltsOperationTrace,omitempty"`

	// 企业项目信息，如果未指定则不开启，默认不开启。
	SysTags *[]SysTagsRes `json:"sys_tags,omitempty"`

	// 创建的图是否启用细粒度权限控制，默认不启用，值为false。如果设置为true，创建的图所有用户都没有权限，需要调用业务面细粒度权限控制API进行授权操作才可以访问图。
	EnableRBAC *bool `json:"enableRBAC,omitempty"`

	// 创建的图是否开启全文索引控制，默认不启用，值为false。如果设置为true，十亿增强版-规格版图支持全文索引，创建图时会创建云搜索服务集群。 >开启全文索引功能。如果CSS服务已经部署，图实例会自动创建CSS集群，图创建时间较长。如果CSS服务没有部署则图创建失败。
	EnableFullTextIndex *bool `json:"enableFullTextIndex,omitempty"`
}

func (o Graph) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Graph struct{}"
	}

	return strings.Join([]string{"Graph", string(data)}, " ")
}
