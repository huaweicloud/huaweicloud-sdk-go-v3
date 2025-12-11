package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VirsubnetCidrReservation struct {

	// **参数解释**： 子网预留网段的资源ID。子网预留网段创建成功后，会生成一个子网预留网段 ID，是子网预留网段对应的唯一标识。 **取值范围**： 带“-”的标准UUID格式。
	Id string `json:"id"`

	// **参数解释**： 子网预留网段所在子网的ID。 **取值范围**： 带“-”的标准UUID格式。
	VirsubnetId string `json:"virsubnet_id"`

	// **参数解释**： 子网预留网段所在VPC的ID。 **取值范围**： 带“-”的标准UUID格式。
	VpcId string `json:"vpc_id"`

	// **参数解释**： 子网预留网段的IP版本。 **取值范围**： - 4：IPv4 - 6：IPv6
	IpVersion int32 `json:"ip_version"`

	// **参数解释**： 子网预留网段的IP网段。 **取值范围**： CIDR格式，掩码长度最小值为“所属子网的网段掩码 + 2”，最大值为32（IPv4）或128（IPv6）。
	Cidr string `json:"cidr"`

	// **参数解释**： 子网预留网段的名称。 **取值范围**： 1-64个字符，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）。
	Name string `json:"name"`

	// **参数解释**： 子网预留网段的描述信息。 **取值范围**： 0-255个字符，不能包含“<”和“>”。
	Description string `json:"description"`

	// **参数解释**： 子网预留网段所属的项目ID。 **取值范围**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 子网预留网段的创建时间。 **取值范围**： UTC时间格式，yyyy-MM-ddTHH:mm:ssZ。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// **参数解释**： 子网预留网段最近一次更新的时间。 **取值范围**： UTC时间格式，yyyy-MM-ddTHH:mm:ssZ。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o VirsubnetCidrReservation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirsubnetCidrReservation struct{}"
	}

	return strings.Join([]string{"VirsubnetCidrReservation", string(data)}, " ")
}
