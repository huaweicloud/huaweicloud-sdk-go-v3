package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTriggersDetailsRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType ListTriggersDetailsRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`
}

func (o ListTriggersDetailsRequest) String() string {
	data, err := json.Marshal(o)
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

func (c ListTriggersDetailsRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListTriggersDetailsRequestContentType) UnmarshalJSON(b []byte) error {
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
