package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVirsubnetCidrReservationsRequest Request Object
type ListVirsubnetCidrReservationsRequest struct {

	// **参数解释**： 每页返回的资源个数。 **取值范围**： 0~2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 分页查询起始的资源ID，为空时查询第一页。 **取值范围**： 子网预留网段的资源ID。
	Marker *string `json:"marker,omitempty"`

	// **参数解释**： 子网预留网段的资源ID。可以使用该字段过滤子网预留网段，支持传入多个ID过滤。 **取值范围**： 不涉及。
	Id *[]string `json:"id,omitempty"`

	// **参数解释**： 子网预留网段所在的子网ID。可以使用该字段过滤子网预留网段，支持传入多个ID过滤。 **取值范围**： 不涉及。
	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`

	// **参数解释**： 子网预留网段的IP网段。可以使用该字段过滤子网预留网段，支持传入多个IP网段过滤。 **取值范围**： 不涉及。
	Cidr *[]string `json:"cidr,omitempty"`

	// **参数解释**： 子网预留网段所在子网的IP版本。可以使用该字段过滤子网预留网段，支持传入多个IP版本过滤。 **取值范围**： - 4：过滤出IPv4子网预留网段。 - 6：过滤出IPv6子网预留网段。
	IpVersion *[]int32 `json:"ip_version,omitempty"`

	// **参数解释**： 子网预留网段的名称。可以使用该字段过滤满足条件的子网预留网段，支持传入多个名称过滤。 **取值范围**： 不涉及。
	Name *[]string `json:"name,omitempty"`

	// **参数解释**： 子网预留网段的描述信息。可以使用该字段过滤子网预留网段，支持传入多个描述信息进行过滤。 **取值范围**： 不涉及。
	Description *[]string `json:"description,omitempty"`

	// **参数解释**： 子网预留网段所属的企业项目ID。可以使用该字段过滤某个企业项目下的子网预留网段。 **取值范围**： - 最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。 - 若需要查询当前用户所有有权限查看企业项目绑定的子网预留网段，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListVirsubnetCidrReservationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirsubnetCidrReservationsRequest struct{}"
	}

	return strings.Join([]string{"ListVirsubnetCidrReservationsRequest", string(data)}, " ")
}
