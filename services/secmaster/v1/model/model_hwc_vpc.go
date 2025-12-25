package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcVpc 云虚拟私有云
type HwcVpc struct {

	// VPC对应的唯一标识
	Id string `json:"id"`

	// VPC对应的名称
	Name string `json:"name"`

	// VPC的描述信息
	Description *string `json:"description,omitempty"`

	// CFW开启或安全组规则配置状态：OPEN | CLOSE
	ProtectedStatus string `json:"protected_status"`

	// VPC下可用子网的范围 取值范围： 10.0.0.0/8~10.255.255.240/28 172.16.0.0/12 ~ 172.31.255.240/28 192.168.0.0/16 ~ 192.168.255.240/28 不指定cidr时，默认值为“” 约束：必须是ipv4 cidr格式，例如:192.168.0.0/16
	Cidr *string `json:"cidr,omitempty"`

	// VPC的扩展网段 约束：目前只支持ipv4
	ExtendCidrs *[]string `json:"extend_cidrs,omitempty"`

	// VPC对应的状态 取值范围： PENDING：创建中 ACTIVE：创建成功
	Status string `json:"status"`

	// VPC所属的项目ID
	ProjectId string `json:"project_id"`

	// VPC所属的企业项目ID。 取值范围：最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// VPC创建时间 取值范围：UTC时间格式，yyyy-MM-ddTHH:mm:ss
	CreatedAt string `json:"created_at"`

	// VPC更新时间 取值范围：UTC时间格式，yyyy-MM-ddTHH:mm:ss
	UpdatedAt string `json:"updated_at"`

	// VPC关联资产类型和数量 取值范围：目前只返回VPC关联的routetable和virsubnet。virsubnet数量为ipv4和ipv6子网总数。
	CloudResources *[]HwcVpcCloudResource `json:"cloud_resources,omitempty"`

	// VPC的标签信息，详情参见Tag对象 取值范围：0-10个标签键值对
	Tags *[]Tag `json:"tags,omitempty"`
}

func (o HwcVpc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcVpc struct{}"
	}

	return strings.Join([]string{"HwcVpc", string(data)}, " ")
}
