package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 查询api版本结构
type ShowApiVersionParams struct {
	// 版本ID（版本号），如v1。

	Id string `json:"id"`
	// 版本号查询链接

	Links []ShowApiVersionLinksParams `json:"links"`
	// 若该版本API支持微版本，则返回支持的最大微版本号，如果不支持微版本，则返回空。

	Version string `json:"version"`
	// 版本状态，为如下3种： CURRENT：表示该版本为主推版本 SUPPORTED：表示为老版本，但是现在还继续支持 DEPRECATED：表示为废弃版本，存在后续删除的可能

	Status ShowApiVersionParamsStatus `json:"status"`
	// 版本发布时间，采用UTC时间表示。如v1发布的时间2018-05-30T15:00:00Z。

	Updated string `json:"updated"`
	// 若该版本API 支持微版本，则返回支持的最小微版本号，如果不支持微版本，则返回空。

	MinVersion string `json:"min_version"`
}

func (o ShowApiVersionParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiVersionParams struct{}"
	}

	return strings.Join([]string{"ShowApiVersionParams", string(data)}, " ")
}

type ShowApiVersionParamsStatus struct {
	value string
}

type ShowApiVersionParamsStatusEnum struct {
	CURRENT    ShowApiVersionParamsStatus
	SUPPORTED  ShowApiVersionParamsStatus
	DEPRECATED ShowApiVersionParamsStatus
}

func GetShowApiVersionParamsStatusEnum() ShowApiVersionParamsStatusEnum {
	return ShowApiVersionParamsStatusEnum{
		CURRENT: ShowApiVersionParamsStatus{
			value: "CURRENT",
		},
		SUPPORTED: ShowApiVersionParamsStatus{
			value: "SUPPORTED",
		},
		DEPRECATED: ShowApiVersionParamsStatus{
			value: "DEPRECATED",
		},
	}
}

func (c ShowApiVersionParamsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiVersionParamsStatus) UnmarshalJSON(b []byte) error {
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
