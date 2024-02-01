package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListNatGatewayDnatRulesRequest Request Object
type ListNatGatewayDnatRulesRequest struct {

	// 解冻/冻结状态。 取值范围： \"true\"：解冻 \"false\"：冻结
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// Floatingip对外提供服务的端口号。 取值范围：0~65535。
	ExternalServicePort *int32 `json:"external_service_port,omitempty"`

	// 弹性公网的IP地址。
	FloatingIpAddress *string `json:"floating_ip_address,omitempty"`

	// DNAT规则的状态。 取值为： \"ACTIVE\": 可用 \"PENDING_CREATE\"：创建中 \"PENDING_UPDATE\"：更新中 \"PENDING_DELETE\"：删除中 \"EIP_FREEZED\"：EIP冻结 \"INACTIVE\"：不可用
	Status *[]ListNatGatewayDnatRulesRequestStatus `json:"status,omitempty"`

	// 弹性公网IP的id。
	FloatingIpId *string `json:"floating_ip_id,omitempty"`

	// 虚拟机或者裸机对外提供服务的协议端口号。 取值范围：0~65535。
	InternalServicePort *int32 `json:"internal_service_port,omitempty"`

	// 功能说明：每页返回的个数。 取值范围：0~2000。 默认值：2000。
	Limit *int32 `json:"limit,omitempty"`

	// DNAT规则的ID。
	Id *string `json:"id,omitempty"`

	// DNAT规则的描述，长度范围小于等于255个字符，不能包含“<”和“>”。
	Description *string `json:"description,omitempty"`

	// DNAT规则的创建时间，格式是yyyy-mm-dd hh:mm:ss.SSSSSS。
	CreatedAt *string `json:"created_at,omitempty"`

	// 公网NAT网关实例的ID。
	NatGatewayId *[]string `json:"nat_gateway_id,omitempty"`

	// 虚拟机或者裸机的Port ID，对应虚拟私有云场景，与private_ip参数二选一。
	PortId *string `json:"port_id,omitempty"`

	// 用户私有IP地址，对应专线、云连接场景，与port_id参数二选一。
	PrivateIp *string `json:"private_ip,omitempty"`

	// 协议类型，目前支持TCP/tcp、UDP/udp、ANY/any。 对应协议号6、17、0。
	Protocol *[]string `json:"protocol,omitempty"`

	// 分页查询的起始资源ID，表示从指定资源的下一条记录开始查询。 - 若不传入marker和limit参数，查询结果返回第一页全部资源记录（默认2000条）。 - 若不传入marker参数，limit为10，查询结果返回第1~10条资源记录。 - 若marker为第10条记录的资源ID，limit为10，查询结果返回第11~20条资源记录。 - 若marker为第10条记录的资源ID，不传入limit参数，查询结果返回第11条及之后的资源记录（默认2000条）。
	Marker *string `json:"marker,omitempty"`
}

func (o ListNatGatewayDnatRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewayDnatRulesRequest struct{}"
	}

	return strings.Join([]string{"ListNatGatewayDnatRulesRequest", string(data)}, " ")
}

type ListNatGatewayDnatRulesRequestStatus struct {
	value string
}

type ListNatGatewayDnatRulesRequestStatusEnum struct {
	ACTIVE         ListNatGatewayDnatRulesRequestStatus
	PENDING_CREATE ListNatGatewayDnatRulesRequestStatus
	PENDING_UPDATE ListNatGatewayDnatRulesRequestStatus
	PENDING_DELETE ListNatGatewayDnatRulesRequestStatus
	EIP_FREEZED    ListNatGatewayDnatRulesRequestStatus
	INACTIVE       ListNatGatewayDnatRulesRequestStatus
}

func GetListNatGatewayDnatRulesRequestStatusEnum() ListNatGatewayDnatRulesRequestStatusEnum {
	return ListNatGatewayDnatRulesRequestStatusEnum{
		ACTIVE: ListNatGatewayDnatRulesRequestStatus{
			value: "ACTIVE",
		},
		PENDING_CREATE: ListNatGatewayDnatRulesRequestStatus{
			value: "PENDING_CREATE",
		},
		PENDING_UPDATE: ListNatGatewayDnatRulesRequestStatus{
			value: "PENDING_UPDATE",
		},
		PENDING_DELETE: ListNatGatewayDnatRulesRequestStatus{
			value: "PENDING_DELETE",
		},
		EIP_FREEZED: ListNatGatewayDnatRulesRequestStatus{
			value: "EIP_FREEZED",
		},
		INACTIVE: ListNatGatewayDnatRulesRequestStatus{
			value: "INACTIVE",
		},
	}
}

func (c ListNatGatewayDnatRulesRequestStatus) Value() string {
	return c.value
}

func (c ListNatGatewayDnatRulesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListNatGatewayDnatRulesRequestStatus) UnmarshalJSON(b []byte) error {
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
