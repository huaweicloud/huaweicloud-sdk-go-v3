package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Virsubnet struct {

	// **参数解释**： 子网的资源ID。子网创建成功后，会生成一个子网 ID，是子网对应的唯一标识。 **取值范围**： 带“-”的标准UUID格式。
	Id string `json:"id"`

	// **参数解释**： 子网的名称。 **取值范围**： 1-64个字符，支持数字、字母、中文字符、_(下划线)、-（中划线）、.（点）。
	Name string `json:"name"`

	// **参数解释**： 子网的描述信息。 **取值范围**： 0-255个字符，不能包含“<”和“>”。
	Description string `json:"description"`

	// **参数解释**： 子网的DNS服务器地址列表。 **取值范围**： 不涉及。
	DnsNameservers []string `json:"dns_nameservers"`

	// **参数解释**： 子网的可用区ID。 **取值范围**： 不涉及。
	ZoneId string `json:"zone_id"`

	// **参数解释**： 子网所属VPC的资源ID。 **取值范围**： 不涉及。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 子网的状态。 **取值范围**： - ACTIVE：表示子网已挂载到VPC上。 - UNKNOWN：表示子网还未挂载到VPC上。 - ERROR：表示子网状态故障。
	Status string `json:"status"`

	// **参数解释**： 子网所属的项目ID。 **取值范围**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 子网的作用域（边缘云场景）。 **取值范围**： - center：表示作用域为中心。 - {publicBorderGroup}：表示作用域为具体的公网边界组。公网边界组限制子网的可用区范围，可关联多个边缘可用区。
	Scope string `json:"scope"`

	// **参数解释**： OpenStack Neutron子网信息。 **取值范围**： 不涉及。
	SubnetCidrs []SubnetCidr `json:"subnet_cidrs"`

	// **参数解释**： 子网的标签信息，包括标签键和标签值，可用来分类和标识资源。详情请参见Tag对象。 **取值范围**： 不涉及。
	Tags []ResponseTag `json:"tags"`

	// **参数解释**： 子网的DHCP属性，支持配置NTP地址、DNS域名或租约到期时间。 **取值范围**： 不涉及。
	ExtraDhcpOpts []SubnetExtraDhcpOpt `json:"extra_dhcp_opts"`

	// **参数解释**： 子网的创建时间。 **取值范围**： UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// **参数解释**： 子网的更新时间。 **取值范围**： UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o Virsubnet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Virsubnet struct{}"
	}

	return strings.Join([]string{"Virsubnet", string(data)}, " ")
}
