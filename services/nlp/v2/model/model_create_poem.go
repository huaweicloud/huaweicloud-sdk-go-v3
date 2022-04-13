package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

//
type CreatePoem struct {
	// 诗歌标题，目前仅支持UTF-8编码，仅支持中文，长度为1-10

	Title string `json:"title"`
	// 诗歌类型，取值如下： 0：五言绝句； 1：七言绝句； 2：五言律诗； 3：七言律诗；

	Type CreatePoemType `json:"type"`
	// 藏头诗，取值如下： 取值为true，为藏头诗； 取值为false，非藏头诗； 默认取值为false。

	Acrostic *bool `json:"acrostic,omitempty"`
}

func (o CreatePoem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePoem struct{}"
	}

	return strings.Join([]string{"CreatePoem", string(data)}, " ")
}

type CreatePoemType struct {
	value int32
}

type CreatePoemTypeEnum struct {
	E_0 CreatePoemType
	E_1 CreatePoemType
	E_2 CreatePoemType
	E_3 CreatePoemType
}

func GetCreatePoemTypeEnum() CreatePoemTypeEnum {
	return CreatePoemTypeEnum{
		E_0: CreatePoemType{
			value: 0,
		}, E_1: CreatePoemType{
			value: 1,
		}, E_2: CreatePoemType{
			value: 2,
		}, E_3: CreatePoemType{
			value: 3,
		},
	}
}

func (c CreatePoemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePoemType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
