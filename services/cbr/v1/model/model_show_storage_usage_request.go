package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowStorageUsageRequest struct {

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移值
	Offset *int32 `json:"offset,omitempty"`

	// 支持按照备份对象ID过滤
	ResourceId *string `json:"resource_id,omitempty"`

	// 支持按照备份对象类型过滤
	ResourceType *ShowStorageUsageRequestResourceType `json:"resource_type,omitempty"`
}

func (o ShowStorageUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStorageUsageRequest struct{}"
	}

	return strings.Join([]string{"ShowStorageUsageRequest", string(data)}, " ")
}

type ShowStorageUsageRequestResourceType struct {
	value string
}

type ShowStorageUsageRequestResourceTypeEnum struct {
	OSNOVASERVER              ShowStorageUsageRequestResourceType
	OSIRONICBARE_METAL_SERVER ShowStorageUsageRequestResourceType
}

func GetShowStorageUsageRequestResourceTypeEnum() ShowStorageUsageRequestResourceTypeEnum {
	return ShowStorageUsageRequestResourceTypeEnum{
		OSNOVASERVER: ShowStorageUsageRequestResourceType{
			value: "OS::Nova::Server",
		},
		OSIRONICBARE_METAL_SERVER: ShowStorageUsageRequestResourceType{
			value: "OS::Ironic::BareMetalServer",
		},
	}
}

func (c ShowStorageUsageRequestResourceType) Value() string {
	return c.value
}

func (c ShowStorageUsageRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowStorageUsageRequestResourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
