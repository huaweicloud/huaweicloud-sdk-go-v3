package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 描述VPCEP服务API版本信息列表
type Version struct {
	// 版本状态。● CURRENT：表示该版本为主推版本。● SUPPORT：表示为老版本，但是现在还在继续支持。● DEPRECATED：表示为废弃版本，存在后续删除的可能。

	Status *VersionStatus `json:"status,omitempty"`
	// 版本ID。

	Id *VersionId `json:"id,omitempty"`
	// 版本发布时间。采用UTC时间格式，格式为：YYYY-MMDDTHH:MM:SSZ

	Updated *string `json:"updated,omitempty"`
	// 支持的版本号。

	Version *string `json:"version,omitempty"`
	// 支持的微版本号。若该版本API不支持微版本，则为空。

	MinVersion *string `json:"min_version,omitempty"`
	// API的URL地址

	Links *[]Link `json:"links,omitempty"`
}

func (o Version) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Version struct{}"
	}

	return strings.Join([]string{"Version", string(data)}, " ")
}

type VersionStatus struct {
	value string
}

type VersionStatusEnum struct {
	CURRENT    VersionStatus
	SUPPORT    VersionStatus
	DEPRECATED VersionStatus
}

func GetVersionStatusEnum() VersionStatusEnum {
	return VersionStatusEnum{
		CURRENT: VersionStatus{
			value: "CURRENT",
		},
		SUPPORT: VersionStatus{
			value: "SUPPORT",
		},
		DEPRECATED: VersionStatus{
			value: "DEPRECATED",
		},
	}
}

func (c VersionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionStatus) UnmarshalJSON(b []byte) error {
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

type VersionId struct {
	value string
}

type VersionIdEnum struct {
	V1 VersionId
}

func GetVersionIdEnum() VersionIdEnum {
	return VersionIdEnum{
		V1: VersionId{
			value: "v1",
		},
	}
}

func (c VersionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionId) UnmarshalJSON(b []byte) error {
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
