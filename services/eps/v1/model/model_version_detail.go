package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/sdktime"
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"
	"errors"
	"strings"
)

// 版本详情
type VersionDetail struct {
	// 版本ID（版本号），如v1.0。

	Id string `json:"id"`
	// API的URL地址。

	Links []Link `json:"links"`
	// 若该版本API支持微版本，则返回支持的最大微版本号，如果不支持微版本，则返回空。

	Version string `json:"version"`
	// 版本状态，为如下3种： CURRENT：表示该版本为主推版本。 SUPPORTED：表示为老版本，但是现在还继续支持。 DEPRECATED：表示为废弃版本，存在后续删除的可能。

	Status VersionDetailStatus `json:"status"`
	// 版本发布时间，采用UTC时间表示。如v1.0发布的时间2016-12-09T00:00:00Z。

	Updated *sdktime.SdkTime `json:"updated"`
	// 若该版本API 支持微版本，则返回支持的最小微版本号， 如果不支持微版本，则返回空。

	MinVersion string `json:"min_version"`
}

func (o VersionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionDetail struct{}"
	}

	return strings.Join([]string{"VersionDetail", string(data)}, " ")
}

type VersionDetailStatus struct {
	value string
}

type VersionDetailStatusEnum struct {
	CURRENT    VersionDetailStatus
	SUPPORTED  VersionDetailStatus
	DEPRECATED VersionDetailStatus
}

func GetVersionDetailStatusEnum() VersionDetailStatusEnum {
	return VersionDetailStatusEnum{
		CURRENT: VersionDetailStatus{
			value: "CURRENT",
		},
		SUPPORTED: VersionDetailStatus{
			value: "SUPPORTED",
		},
		DEPRECATED: VersionDetailStatus{
			value: "DEPRECATED",
		},
	}
}

func (c VersionDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VersionDetailStatus) UnmarshalJSON(b []byte) error {
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
