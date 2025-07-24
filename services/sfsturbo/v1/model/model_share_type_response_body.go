package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShareTypeResponseBody struct {

	// 文件系统类型
	ShareType *string `json:"share_type,omitempty"`

	// 文件系统场景
	Scenario *string `json:"scenario,omitempty"`

	Attribution *ShareTypesAttribution `json:"attribution,omitempty"`

	// 文件系统规格ID
	SpecId *string `json:"spec_id,omitempty"`

	ShareTypeUsage *ShareTypeUsage `json:"share_type_usage,omitempty"`

	// 是否支持包周期
	SupportPeriod *bool `json:"support_period,omitempty"`

	// 可用区
	AvailableZones *[]ShareTypeAvailableZone `json:"available_zones,omitempty"`

	// 文件系统规格编码
	SpecCode *string `json:"spec_code,omitempty"`

	// 存储类型
	StorageMedia *ShareTypeResponseBodyStorageMedia `json:"storage_media,omitempty"`

	// 实例支持的特性列表
	Features *[]string `json:"features,omitempty"`
}

func (o ShareTypeResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShareTypeResponseBody struct{}"
	}

	return strings.Join([]string{"ShareTypeResponseBody", string(data)}, " ")
}

type ShareTypeResponseBodyStorageMedia struct {
	value string
}

type ShareTypeResponseBodyStorageMediaEnum struct {
	HDD  ShareTypeResponseBodyStorageMedia
	SSD  ShareTypeResponseBodyStorageMedia
	ESSD ShareTypeResponseBodyStorageMedia
}

func GetShareTypeResponseBodyStorageMediaEnum() ShareTypeResponseBodyStorageMediaEnum {
	return ShareTypeResponseBodyStorageMediaEnum{
		HDD: ShareTypeResponseBodyStorageMedia{
			value: "HDD",
		},
		SSD: ShareTypeResponseBodyStorageMedia{
			value: "SSD",
		},
		ESSD: ShareTypeResponseBodyStorageMedia{
			value: "ESSD",
		},
	}
}

func (c ShareTypeResponseBodyStorageMedia) Value() string {
	return c.value
}

func (c ShareTypeResponseBodyStorageMedia) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShareTypeResponseBodyStorageMedia) UnmarshalJSON(b []byte) error {
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
