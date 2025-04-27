package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ReqAllocateDeh 创建专属主机消息体。
type ReqAllocateDeh struct {

	// 专属主机名称。
	Name string `json:"name"`

	// 在创建云服务器时（未指定专属主机ID），是否允许云服务器自动分配在一台可用的专属主机上。 取值范围：“on”或“off”。 默认值：“on”。
	AutoPlacement *ReqAllocateDehAutoPlacement `json:"auto_placement,omitempty"`

	// 专属主机所属AZ。
	AvailabilityZone string `json:"availability_zone"`

	// 专属主机类型。
	HostType string `json:"host_type"`

	// 待分配的专属主机数量。
	Quantity int32 `json:"quantity"`

	// 专属主机标签列表。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	ExtendParam *ReqAllocateDehExtendParam `json:"extend_param,omitempty"`
}

func (o ReqAllocateDeh) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqAllocateDeh struct{}"
	}

	return strings.Join([]string{"ReqAllocateDeh", string(data)}, " ")
}

type ReqAllocateDehAutoPlacement struct {
	value string
}

type ReqAllocateDehAutoPlacementEnum struct {
	OFF ReqAllocateDehAutoPlacement
	ON  ReqAllocateDehAutoPlacement
}

func GetReqAllocateDehAutoPlacementEnum() ReqAllocateDehAutoPlacementEnum {
	return ReqAllocateDehAutoPlacementEnum{
		OFF: ReqAllocateDehAutoPlacement{
			value: "off",
		},
		ON: ReqAllocateDehAutoPlacement{
			value: "on",
		},
	}
}

func (c ReqAllocateDehAutoPlacement) Value() string {
	return c.value
}

func (c ReqAllocateDehAutoPlacement) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqAllocateDehAutoPlacement) UnmarshalJSON(b []byte) error {
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
