package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type DuplicateApiInfo struct {
	// API ID

	ApiId *string `json:"api_id,omitempty"`
	// API名称

	ApiName *string `json:"api_name,omitempty"`
	// 分组名称

	GroupName *string `json:"group_name,omitempty"`
	// 分组ID

	GroupId *string `json:"group_id,omitempty"`
	// API描述

	Remark *string `json:"remark,omitempty"`
	// api类型： - self-owned：自有API - authorized：授权API

	ApiType *DuplicateApiInfoApiType `json:"api_type,omitempty"`
}

func (o DuplicateApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DuplicateApiInfo struct{}"
	}

	return strings.Join([]string{"DuplicateApiInfo", string(data)}, " ")
}

type DuplicateApiInfoApiType struct {
	value string
}

type DuplicateApiInfoApiTypeEnum struct {
	SELF_OWNED DuplicateApiInfoApiType
	AUTHORIZED DuplicateApiInfoApiType
}

func GetDuplicateApiInfoApiTypeEnum() DuplicateApiInfoApiTypeEnum {
	return DuplicateApiInfoApiTypeEnum{
		SELF_OWNED: DuplicateApiInfoApiType{
			value: "self-owned",
		},
		AUTHORIZED: DuplicateApiInfoApiType{
			value: "authorized",
		},
	}
}

func (c DuplicateApiInfoApiType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DuplicateApiInfoApiType) UnmarshalJSON(b []byte) error {
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
