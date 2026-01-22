package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRepoAccessoriesRequest Request Object
type ListRepoAccessoriesRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListRepoAccessoriesRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace"`

	// 镜像仓库名称
	Repository string `json:"repository"`

	// 镜像版本
	Tag string `json:"tag"`

	// 返回条数。如果不传该参数默认返回10条记录，最大支持100条记录
	Limit *int32 `json:"limit,omitempty"`

	// 起始索引,默认值为0
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListRepoAccessoriesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepoAccessoriesRequest struct{}"
	}

	return strings.Join([]string{"ListRepoAccessoriesRequest", string(data)}, " ")
}

type ListRepoAccessoriesRequestContentType struct {
	value string
}

type ListRepoAccessoriesRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListRepoAccessoriesRequestContentType
	APPLICATION_JSON             ListRepoAccessoriesRequestContentType
}

func GetListRepoAccessoriesRequestContentTypeEnum() ListRepoAccessoriesRequestContentTypeEnum {
	return ListRepoAccessoriesRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListRepoAccessoriesRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListRepoAccessoriesRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListRepoAccessoriesRequestContentType) Value() string {
	return c.value
}

func (c ListRepoAccessoriesRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepoAccessoriesRequestContentType) UnmarshalJSON(b []byte) error {
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
