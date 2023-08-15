package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTriggersDetailsRequest Request Object
type ListTriggersDetailsRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListTriggersDetailsRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace"`

	// 镜像仓库名称
	Repository string `json:"repository"`
}

func (o ListTriggersDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTriggersDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListTriggersDetailsRequest", string(data)}, " ")
}

type ListTriggersDetailsRequestContentType struct {
	value string
}

type ListTriggersDetailsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListTriggersDetailsRequestContentType
	APPLICATION_JSON             ListTriggersDetailsRequestContentType
}

func GetListTriggersDetailsRequestContentTypeEnum() ListTriggersDetailsRequestContentTypeEnum {
	return ListTriggersDetailsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListTriggersDetailsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListTriggersDetailsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListTriggersDetailsRequestContentType) Value() string {
	return c.value
}

func (c ListTriggersDetailsRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTriggersDetailsRequestContentType) UnmarshalJSON(b []byte) error {
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
