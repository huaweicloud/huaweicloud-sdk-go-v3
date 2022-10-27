package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListVolumesRequest struct {

	// 环境id。
	XEnvironmentID string `json:"X-Environment-ID"`

	// 租户的企业项目id。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 资源类型。
	ResourceType ListVolumesRequestResourceType `json:"resource_type"`
}

func (o ListVolumesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesRequest struct{}"
	}

	return strings.Join([]string{"ListVolumesRequest", string(data)}, " ")
}

type ListVolumesRequestResourceType struct {
	value string
}

type ListVolumesRequestResourceTypeEnum struct {
	OBS ListVolumesRequestResourceType
}

func GetListVolumesRequestResourceTypeEnum() ListVolumesRequestResourceTypeEnum {
	return ListVolumesRequestResourceTypeEnum{
		OBS: ListVolumesRequestResourceType{
			value: "obs",
		},
	}
}

func (c ListVolumesRequestResourceType) Value() string {
	return c.value
}

func (c ListVolumesRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVolumesRequestResourceType) UnmarshalJSON(b []byte) error {
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
