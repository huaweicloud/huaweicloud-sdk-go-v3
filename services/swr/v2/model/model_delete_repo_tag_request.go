package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type DeleteRepoTagRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType DeleteRepoTagRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`
	// 镜像版本名称

	Tag string `json:"tag"`
}

func (o DeleteRepoTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRepoTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteRepoTagRequest", string(data)}, " ")
}

type DeleteRepoTagRequestContentType struct {
	value string
}

type DeleteRepoTagRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 DeleteRepoTagRequestContentType
	APPLICATION_JSON             DeleteRepoTagRequestContentType
}

func GetDeleteRepoTagRequestContentTypeEnum() DeleteRepoTagRequestContentTypeEnum {
	return DeleteRepoTagRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: DeleteRepoTagRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: DeleteRepoTagRequestContentType{
			value: "application/json",
		},
	}
}

func (c DeleteRepoTagRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DeleteRepoTagRequestContentType) UnmarshalJSON(b []byte) error {
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
