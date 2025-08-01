package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRepoTagRequest Request Object
type ShowRepoTagRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ShowRepoTagRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace"`

	// 镜像仓库名称
	Repository string `json:"repository"`

	// 镜像版本名称
	Tag string `json:"tag"`
}

func (o ShowRepoTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoTagRequest struct{}"
	}

	return strings.Join([]string{"ShowRepoTagRequest", string(data)}, " ")
}

type ShowRepoTagRequestContentType struct {
	value string
}

type ShowRepoTagRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowRepoTagRequestContentType
	APPLICATION_JSON             ShowRepoTagRequestContentType
}

func GetShowRepoTagRequestContentTypeEnum() ShowRepoTagRequestContentTypeEnum {
	return ShowRepoTagRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowRepoTagRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowRepoTagRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowRepoTagRequestContentType) Value() string {
	return c.value
}

func (c ShowRepoTagRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepoTagRequestContentType) UnmarshalJSON(b []byte) error {
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
