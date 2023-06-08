package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListAvailableZonesRequest struct {

	// 区域id，例如cn-north-4
	RegionId string `json:"region_id"`

	// 可用区范围
	Scope *ListAvailableZonesRequestScope `json:"scope,omitempty"`
}

func (o ListAvailableZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesRequest struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesRequest", string(data)}, " ")
}

type ListAvailableZonesRequestScope struct {
	value string
}

type ListAvailableZonesRequestScopeEnum struct {
	CENTER ListAvailableZonesRequestScope
	EDGE   ListAvailableZonesRequestScope
}

func GetListAvailableZonesRequestScopeEnum() ListAvailableZonesRequestScopeEnum {
	return ListAvailableZonesRequestScopeEnum{
		CENTER: ListAvailableZonesRequestScope{
			value: "Center",
		},
		EDGE: ListAvailableZonesRequestScope{
			value: "Edge",
		},
	}
}

func (c ListAvailableZonesRequestScope) Value() string {
	return c.value
}

func (c ListAvailableZonesRequestScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAvailableZonesRequestScope) UnmarshalJSON(b []byte) error {
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
