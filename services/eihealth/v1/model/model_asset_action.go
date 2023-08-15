package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

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
