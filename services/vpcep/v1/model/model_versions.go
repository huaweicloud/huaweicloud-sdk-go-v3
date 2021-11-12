package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// 描述VPCEP服务API版本信息列表
type Versions struct {
	// 版本状态。● CURRENT：表示该版本为主推版本。● SUPPORT：表示为老版本，但是现在还在继续支持。● DEPRECATED：表示为废弃版本，存在后续删除的可能。

	Status *VersionsStatus `json:"status,omitempty"`
	// 版本ID。

	Id *VersionsId `json:"id,omitempty"`
	// 版本发布时间。采用UTC时间格式，格式为：YYYY-MMDDTHH:MM:SSZ

	Updated *sdktime.SdkTime `json:"updated,omitempty"`
	// 支持的版本号。

	Version *string `json:"version,omitempty"`
	// 支持的微版本号。若该版本API不支持微版本，则为空。

	MinVersion *string `json:"min_version,omitempty"`
	// API的URL地址

	Links *[]Links `json:"links,omitempty"`
}

func (o Versions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Versions struct{}"
	}

	return strings.Join([]string{"Versions", string(data)}, " ")
}

type VersionsStatus struct {
	value string
}

type VersionsStatusEnum struct {
	CURRENT    VersionsStatus
	SUPPORT    VersionsStatus
	DEPRECATED VersionsStatus
}

func GetVersionsStatusEnum() VersionsStatusEnum {
	return VersionsStatusEnum{
		CURRENT: VersionsStatus{
			value: "CURRENT",
		},
		SUPPORT: VersionsStatus{
			value: "SUPPORT",
		},
		DEPRECATED: VersionsStatus{
			value: "DEPRECATED",
		},
	}
}

func (c VersionsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionsStatus) UnmarshalJSON(b []byte) error {
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

type VersionsId struct {
	value string
}

type VersionsIdEnum struct {
	V1 VersionsId
}

func GetVersionsIdEnum() VersionsIdEnum {
	return VersionsIdEnum{
		V1: VersionsId{
			value: "v1",
		},
	}
}

func (c VersionsId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionsId) UnmarshalJSON(b []byte) error {
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
