package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 版本信息描述
type VersionDetails struct {

	// 版本ID。
	Id string `json:"id" xml:"id"`

	// 版本详情的URL地址。
	Links string `json:"links" xml:"links"`

	// 该版本API的微版本信息。
	Version string `json:"version" xml:"version"`

	// 版本的状态。
	Status VersionDetailsStatus `json:"status" xml:"status"`

	// 版本更新时间。
	Updated string `json:"updated" xml:"updated"`
}

func (o VersionDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionDetails struct{}"
	}

	return strings.Join([]string{"VersionDetails", string(data)}, " ")
}

type VersionDetailsStatus struct {
	value string
}

type VersionDetailsStatusEnum struct {
	CURRENT    VersionDetailsStatus
	SUPPORTED  VersionDetailsStatus
	DEPRECATED VersionDetailsStatus
}

func GetVersionDetailsStatusEnum() VersionDetailsStatusEnum {
	return VersionDetailsStatusEnum{
		CURRENT: VersionDetailsStatus{
			value: "CURRENT",
		},
		SUPPORTED: VersionDetailsStatus{
			value: "SUPPORTED",
		},
		DEPRECATED: VersionDetailsStatus{
			value: "DEPRECATED",
		},
	}
}

func (c VersionDetailsStatus) Value() string {
	return c.value
}

func (c VersionDetailsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionDetailsStatus) UnmarshalJSON(b []byte) error {
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
