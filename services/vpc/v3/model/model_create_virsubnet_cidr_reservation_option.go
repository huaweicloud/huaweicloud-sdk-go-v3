package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVirsubnetCidrReservationOption struct {

	// **参数解释**： 子网预留网段所属的子网ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	VirsubnetId string `json:"virsubnet_id"`

	// **参数解释**： 子网预留网段的IP版本，支持IPv4和IPv6。 **约束限制**： 不涉及。 **取值范围**： - 4：表示IPv4。 - 6：表示IPv6。 **默认取值**： 不涉及。
	IpVersion int32 `json:"ip_version"`

	// **参数解释**： 子网预留网段的IP网段。 **约束限制**： - CIDR格式，掩码长度最小值为“所属子网的网段掩码 + 2”，最大值为32（IPv4）或128（IPv6）。 - cidr和mask参数必须输入一个，两者同时输入时不能冲突。 - 子网预留网段不能包含所属子网的已使用的地址和系统预留地址（子网的第1个和最后2个地址）。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Cidr *string `json:"cidr,omitempty"`

	// **参数解释**： 子网预留网段的IP网段掩码长度。 **约束限制**： - 整数，预留网段的掩码长度最小值为“所属子网的网段掩码 + 2”，最大值为32（IPv4）或128（IPv6）。 - cidr和mask参数必须输入一个，两者同时输入时不能冲突。 - 子网预留网段不能包含所属子网的已使用的地址和系统预留地址（子网的第1个和最后2个地址）。 - 指定掩码长度创建预留网段，最后mask与子网掩码的差值长度的bit位由系统自动分配，例如子网cidr为192.168.21.0/24，子网掩码长度24，指定预留网段长度为27，差值长度27 - 24 = 3，即第25,26,27这3个bit位由系统自动分配。例如：   - 第25-27的bit位分配为011，最终创建出的子网预留网段cidr是192.168.21.96/27，其中96转为二进制是0110 0000；   - 第25-27的bit位分配为110，最终创建出的子网预留网段cidr是192.168.21.192/27，其中192转为二进制是1100 0000。  **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Mask *int32 `json:"mask,omitempty"`

	// **参数解释**： 子网预留网段的名称。 **约束限制**： 1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 子网预留网段的描述信息。 **约束限制**： 0-255个字符，不能包含“<”和“>”。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o CreateVirsubnetCidrReservationOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVirsubnetCidrReservationOption struct{}"
	}

	return strings.Join([]string{"CreateVirsubnetCidrReservationOption", string(data)}, " ")
}
