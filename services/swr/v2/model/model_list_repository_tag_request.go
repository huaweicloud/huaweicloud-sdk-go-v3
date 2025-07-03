package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRepositoryTagRequest Request Object
type ListRepositoryTagRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListRepositoryTagRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace"`

	// 镜像仓库名称
	Repository string `json:"repository"`

	// 返回条数,默认返回100条，最多返回1000条数据。
	Limit *int32 `json:"limit,omitempty"`

	// Start position of the cursor for querying the next page in pagination query.
	Marker *string `json:"marker,omitempty"`

	// 镜像版本名。
	Tag *string `json:"tag,omitempty"`

	// 是否返回镜像的manifest信息
	WithManifest *bool `json:"with_manifest,omitempty"`
}

func (o ListRepositoryTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryTagRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryTagRequest", string(data)}, " ")
}

type ListRepositoryTagRequestContentType struct {
	value string
}

type ListRepositoryTagRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListRepositoryTagRequestContentType
	APPLICATION_JSON             ListRepositoryTagRequestContentType
}

func GetListRepositoryTagRequestContentTypeEnum() ListRepositoryTagRequestContentTypeEnum {
	return ListRepositoryTagRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListRepositoryTagRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListRepositoryTagRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListRepositoryTagRequestContentType) Value() string {
	return c.value
}

func (c ListRepositoryTagRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryTagRequestContentType) UnmarshalJSON(b []byte) error {
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
