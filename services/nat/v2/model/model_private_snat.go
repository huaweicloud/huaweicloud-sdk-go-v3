package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// PrivateSnat SNAT规则的响应体。
type PrivateSnat struct {

	// SNAT规则的ID。
	Id *string `json:"id,omitempty"`

	// 项目的ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 私网NAT网关实例的ID。
	GatewayId *string `json:"gateway_id,omitempty"`

	// 功能说明：规则匹配的CIDR。 取值约束： - 与virsubnet_id参数二选一。 - cidr不能与已有snat规则的网段相同。
	Cidr *string `json:"cidr,omitempty"`

	// 功能说明：规则匹配的子网的ID。 取值约束：与cidr参数二选一。
	VirsubnetId *string `json:"virsubnet_id,omitempty"`

	// SNAT规则的描述。长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`

	// 关联的中转IP详情列表。
	TransitIpAssociations *[]AssociatedTransitIp `json:"transit_ip_associations,omitempty"`

	// SNAT规则的创建时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// SNAT规则的更新时间，遵循UTC时间，格式是yyyy-mm-ddThh:mm:ssZ。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 私网NAT的SNAT规则状态。 取值为： \"ACTIVE\"：正常运行 \"FROZEN\"：冻结 \"INACTIVE\"：不可用
	Status *PrivateSnatStatus `json:"status,omitempty"`
}

func (o PrivateSnat) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrivateSnat struct{}"
	}

	return strings.Join([]string{"PrivateSnat", string(data)}, " ")
}

type PrivateSnatStatus struct {
	value string
}

type PrivateSnatStatusEnum struct {
	ACTIVE   PrivateSnatStatus
	FROZEN   PrivateSnatStatus
	INACTIVE PrivateSnatStatus
}

func GetPrivateSnatStatusEnum() PrivateSnatStatusEnum {
	return PrivateSnatStatusEnum{
		ACTIVE: PrivateSnatStatus{
			value: "ACTIVE",
		},
		FROZEN: PrivateSnatStatus{
			value: "FROZEN",
		},
		INACTIVE: PrivateSnatStatus{
			value: "INACTIVE",
		},
	}
}

func (c PrivateSnatStatus) Value() string {
	return c.value
}

func (c PrivateSnatStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrivateSnatStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
