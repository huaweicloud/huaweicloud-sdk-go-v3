package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DataObjectNetworkList struct {

	// 方向，取值范围：IN | OUT
	Direction *DataObjectNetworkListDirection `json:"direction,omitempty"`

	// 协议，包含7层和4层的协议 参考：IANA registered name https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	Protocol *string `json:"protocol,omitempty"`

	// 源IP地址
	SrcIp *string `json:"src_ip,omitempty"`

	// 源端口，0–65535
	SrcPort *int32 `json:"src_port,omitempty"`

	// 源域名
	SrcDomain *string `json:"src_domain,omitempty"`

	SrcGeo *AlertSrcGeo `json:"src_geo,omitempty"`

	// 目的IP地址
	DestIp *string `json:"dest_ip,omitempty"`

	// 目的端口，0–65535
	DestPort *string `json:"dest_port,omitempty"`

	// 目的域名
	DestDomain *string `json:"dest_domain,omitempty"`

	DestGeo *AlertDestGeo `json:"dest_geo,omitempty"`
}

func (o DataObjectNetworkList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataObjectNetworkList struct{}"
	}

	return strings.Join([]string{"DataObjectNetworkList", string(data)}, " ")
}

type DataObjectNetworkListDirection struct {
	value string
}

type DataObjectNetworkListDirectionEnum struct {
	IN  DataObjectNetworkListDirection
	OUT DataObjectNetworkListDirection
}

func GetDataObjectNetworkListDirectionEnum() DataObjectNetworkListDirectionEnum {
	return DataObjectNetworkListDirectionEnum{
		IN: DataObjectNetworkListDirection{
			value: "IN",
		},
		OUT: DataObjectNetworkListDirection{
			value: "OUT",
		},
	}
}

func (c DataObjectNetworkListDirection) Value() string {
	return c.value
}

func (c DataObjectNetworkListDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataObjectNetworkListDirection) UnmarshalJSON(b []byte) error {
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
