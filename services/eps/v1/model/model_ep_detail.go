package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// EpDetail 企业项目详情
type EpDetail struct {

	// 企业项目ID
	Id string `json:"id"`

	// 企业项目名称
	Name string `json:"name"`

	// 企业项目描述
	Description string `json:"description"`

	// 企业项目状态。1启用，2停用
	Status int32 `json:"status"`

	// 创建时间，格式为UTC格式。如：2018-05-18T06:49:06Z。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 修改时间，格式为UTC格式。如：2018-05-28T02:21:36Z。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`

	// 项目类型： - prod：商用项目 - poc：测试项目
	Type EpDetailType `json:"type"`

	// 删除标识，false-未删除；true-已删除
	DeleteFlag bool `json:"delete_flag"`
}

func (o EpDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EpDetail struct{}"
	}

	return strings.Join([]string{"EpDetail", string(data)}, " ")
}

type EpDetailType struct {
	value string
}

type EpDetailTypeEnum struct {
	PROD EpDetailType
	POC  EpDetailType
}

func GetEpDetailTypeEnum() EpDetailTypeEnum {
	return EpDetailTypeEnum{
		PROD: EpDetailType{
			value: "prod",
		},
		POC: EpDetailType{
			value: "poc",
		},
	}
}

func (c EpDetailType) Value() string {
	return c.value
}

func (c EpDetailType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EpDetailType) UnmarshalJSON(b []byte) error {
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
