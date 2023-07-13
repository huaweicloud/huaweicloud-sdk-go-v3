package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteRepoTagRequest Request Object
type DeleteRepoTagRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType DeleteRepoTagRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace"`

	// 镜像仓库名称
	Repository string `json:"repository"`

	// 镜像版本名称
	Tag string `json:"tag"`
}

func (o DeleteRepoTagRequest) String() string {
	data, err := utils.Marshal(o)
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

func (c DeleteRepoTagRequestContentType) Value() string {
	return c.value
}

func (c DeleteRepoTagRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteRepoTagRequestContentType) UnmarshalJSON(b []byte) error {
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
