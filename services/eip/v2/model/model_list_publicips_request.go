package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPublicipsRequest Request Object
type ListPublicipsRequest struct {

	// 取值为上一页数据的最后一条记录的id，为空时为查询第一页
	Marker *string `json:"marker,omitempty"`

	// 功能说明：每页返回的个数  取值范围：0~intmax
	Limit *int32 `json:"limit,omitempty"`

	// IP地址版本信息  4：IPv4  6：IPv6
	IpVersion *ListPublicipsRequestIpVersion `json:"ip_version,omitempty"`

	// 功能说明：企业项目ID。可以使用该字段过滤某个企业项目下的弹性IP弹性公网IP。  取值范围：最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。“0”表示默认企业项目。若需要查询当前用户所有企业项目绑定的弹性公网IP，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 绑定弹性公网IP的端口id
	PortId *[]string `json:"port_id,omitempty"`

	// IPv4时是申请到的弹性公网IP地址，IPv6时是IPv6地址对应的IPv4地址
	PublicIpAddress *[]string `json:"public_ip_address,omitempty"`

	// 关联端口的私有IP地址
	PrivateIpAddress *[]string `json:"private_ip_address,omitempty"`

	// 弹性公网IP唯一标识
	Id *[]string `json:"id,omitempty"`

	// 共享带宽类型，根据任一共享带宽类型过滤EIP列表。 可以指定多个带宽类型，不同的带宽类型间用逗号分隔。
	AllowShareBandwidthTypeAny *[]string `json:"allow_share_bandwidth_type_any,omitempty"`
}

func (o ListPublicipsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipsRequest struct{}"
	}

	return strings.Join([]string{"ListPublicipsRequest", string(data)}, " ")
}

type ListPublicipsRequestIpVersion struct {
	value int32
}

type ListPublicipsRequestIpVersionEnum struct {
	E_4 ListPublicipsRequestIpVersion
	E_6 ListPublicipsRequestIpVersion
}

func GetListPublicipsRequestIpVersionEnum() ListPublicipsRequestIpVersionEnum {
	return ListPublicipsRequestIpVersionEnum{
		E_4: ListPublicipsRequestIpVersion{
			value: 4,
		}, E_6: ListPublicipsRequestIpVersion{
			value: 6,
		},
	}
}

func (c ListPublicipsRequestIpVersion) Value() int32 {
	return c.value
}

func (c ListPublicipsRequestIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPublicipsRequestIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
