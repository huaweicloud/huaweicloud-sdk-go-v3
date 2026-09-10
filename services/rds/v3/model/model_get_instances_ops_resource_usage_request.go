package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetInstancesOpsResourceUsageRequest Request Object
type GetInstancesOpsResourceUsageRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 资源类型。取值范围：cpu、mem、disk、disk_week、io。
	ResourceType *GetInstancesOpsResourceUsageRequestResourceType `json:"resource_type,omitempty"`
}

func (o GetInstancesOpsResourceUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetInstancesOpsResourceUsageRequest struct{}"
	}

	return strings.Join([]string{"GetInstancesOpsResourceUsageRequest", string(data)}, " ")
}

type GetInstancesOpsResourceUsageRequestResourceType struct {
	value string
}

type GetInstancesOpsResourceUsageRequestResourceTypeEnum struct {
	CPU       GetInstancesOpsResourceUsageRequestResourceType
	MEM       GetInstancesOpsResourceUsageRequestResourceType
	DISK      GetInstancesOpsResourceUsageRequestResourceType
	DISK_WEEK GetInstancesOpsResourceUsageRequestResourceType
	IO        GetInstancesOpsResourceUsageRequestResourceType
}

func GetGetInstancesOpsResourceUsageRequestResourceTypeEnum() GetInstancesOpsResourceUsageRequestResourceTypeEnum {
	return GetInstancesOpsResourceUsageRequestResourceTypeEnum{
		CPU: GetInstancesOpsResourceUsageRequestResourceType{
			value: "cpu",
		},
		MEM: GetInstancesOpsResourceUsageRequestResourceType{
			value: "mem",
		},
		DISK: GetInstancesOpsResourceUsageRequestResourceType{
			value: "disk",
		},
		DISK_WEEK: GetInstancesOpsResourceUsageRequestResourceType{
			value: "disk_week",
		},
		IO: GetInstancesOpsResourceUsageRequestResourceType{
			value: "io",
		},
	}
}

func (c GetInstancesOpsResourceUsageRequestResourceType) Value() string {
	return c.value
}

func (c GetInstancesOpsResourceUsageRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetInstancesOpsResourceUsageRequestResourceType) UnmarshalJSON(b []byte) error {
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
