package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AssetAction struct {
	value string
}

type AssetActionEnum struct {
	RETRY   AssetAction
	CANCEL  AssetAction
	OFFLINE AssetAction
}

func GetAssetActionEnum() AssetActionEnum {
	return AssetActionEnum{
		RETRY: AssetAction{
			value: "RETRY",
		},
		CANCEL: AssetAction{
			value: "CANCEL",
		},
		OFFLINE: AssetAction{
			value: "OFFLINE",
		},
	}
}

func (c AssetAction) Value() string {
	return c.value
}

func (c AssetAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssetAction) UnmarshalJSON(b []byte) error {
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
