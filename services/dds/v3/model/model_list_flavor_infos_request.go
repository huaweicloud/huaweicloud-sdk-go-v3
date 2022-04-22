package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListFlavorInfosRequest struct {

	// 数据库版本类型。取值为“DDS-Community”。
	EngineName *ListFlavorInfosRequestEngineName `json:"engine_name,omitempty"`

	// 数据库版本号。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 索引位置，偏移量。   - 从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询）。   - 取值必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询规格信息上限值。   - 取值范围: 1~100。   - 不传该参数时，默认查询前100条规格信息。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFlavorInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFlavorInfosRequest struct{}"
	}

	return strings.Join([]string{"ListFlavorInfosRequest", string(data)}, " ")
}

type ListFlavorInfosRequestEngineName struct {
	value string
}

type ListFlavorInfosRequestEngineNameEnum struct {
	DDS_COMMUNITY ListFlavorInfosRequestEngineName
	DDS_ENHANCED  ListFlavorInfosRequestEngineName
}

func GetListFlavorInfosRequestEngineNameEnum() ListFlavorInfosRequestEngineNameEnum {
	return ListFlavorInfosRequestEngineNameEnum{
		DDS_COMMUNITY: ListFlavorInfosRequestEngineName{
			value: "DDS-Community",
		},
		DDS_ENHANCED: ListFlavorInfosRequestEngineName{
			value: "DDS-Enhanced",
		},
	}
}

func (c ListFlavorInfosRequestEngineName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFlavorInfosRequestEngineName) UnmarshalJSON(b []byte) error {
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
