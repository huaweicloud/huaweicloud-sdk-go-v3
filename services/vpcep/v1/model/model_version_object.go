package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VersionObject 描述VPCEP服务API版本信息列表
type VersionObject struct {

	// 版本状态。 ● CURRENT：表示该版本为主推版本。 ● SUPPORT：表示为老版本，但是现在还在继续支持。 ● DEPRECATED：表示为废弃版本，存在后续删除的可能。
	Status *VersionObjectStatus `json:"status,omitempty"`

	// 版本ID。
	Id *VersionObjectId `json:"id,omitempty"`

	// 版本发布时间。采用UTC时间格式，格式为：YYYY-MMDDTHH:MM:SSZ
	Updated *string `json:"updated,omitempty"`

	// 支持的版本号。
	Version *string `json:"version,omitempty"`

	// 支持的微版本号。若该版本API不支持微版本，则为空。
	MinVersion *string `json:"min_version,omitempty"`

	// API的URL地址
	Links *[]Link `json:"links,omitempty"`
}

func (o VersionObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionObject struct{}"
	}

	return strings.Join([]string{"VersionObject", string(data)}, " ")
}

type VersionObjectStatus struct {
	value string
}

type VersionObjectStatusEnum struct {
	CURRENT    VersionObjectStatus
	SUPPORT    VersionObjectStatus
	DEPRECATED VersionObjectStatus
}

func GetVersionObjectStatusEnum() VersionObjectStatusEnum {
	return VersionObjectStatusEnum{
		CURRENT: VersionObjectStatus{
			value: "CURRENT",
		},
		SUPPORT: VersionObjectStatus{
			value: "SUPPORT",
		},
		DEPRECATED: VersionObjectStatus{
			value: "DEPRECATED",
		},
	}
}

func (c VersionObjectStatus) Value() string {
	return c.value
}

func (c VersionObjectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionObjectStatus) UnmarshalJSON(b []byte) error {
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

type VersionObjectId struct {
	value string
}

type VersionObjectIdEnum struct {
	V1 VersionObjectId
}

func GetVersionObjectIdEnum() VersionObjectIdEnum {
	return VersionObjectIdEnum{
		V1: VersionObjectId{
			value: "v1",
		},
	}
}

func (c VersionObjectId) Value() string {
	return c.value
}

func (c VersionObjectId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionObjectId) UnmarshalJSON(b []byte) error {
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
