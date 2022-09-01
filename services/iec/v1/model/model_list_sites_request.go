package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListSitesRequest struct {

	// 查询返回边缘站点列表当前页面的数量。 取值范围：0~1000。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询的偏移量。默认为0。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询条件，站点ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 边缘实例所在大区。   大小写通用，皆支持。 支持多个查询，中间使用','分隔。
	Area *string `json:"area,omitempty" xml:"area"`

	// 边缘实例所在省份。  大小写通用，皆支持。 支持多个查询，中间使用“,”分隔。
	Province *string `json:"province,omitempty" xml:"province"`

	// 边缘实例所在城市。  大小写通用，皆支持。 支持多个查询，中间使用“,”分隔。
	City *string `json:"city,omitempty" xml:"city"`

	// 边缘实例规格。
	Flavor *string `json:"flavor,omitempty" xml:"flavor"`

	// 过滤支持磁盘类型的站点，多个类型之间用“,”分割。
	VolumeType *ListSitesRequestVolumeType `json:"volume_type,omitempty" xml:"volume_type"`
}

func (o ListSitesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSitesRequest struct{}"
	}

	return strings.Join([]string{"ListSitesRequest", string(data)}, " ")
}

type ListSitesRequestVolumeType struct {
	value string
}

type ListSitesRequestVolumeTypeEnum struct {
	SATA ListSitesRequestVolumeType
	SAS  ListSitesRequestVolumeType
}

func GetListSitesRequestVolumeTypeEnum() ListSitesRequestVolumeTypeEnum {
	return ListSitesRequestVolumeTypeEnum{
		SATA: ListSitesRequestVolumeType{
			value: "SATA",
		},
		SAS: ListSitesRequestVolumeType{
			value: "SAS",
		},
	}
}

func (c ListSitesRequestVolumeType) Value() string {
	return c.value
}

func (c ListSitesRequestVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSitesRequestVolumeType) UnmarshalJSON(b []byte) error {
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
