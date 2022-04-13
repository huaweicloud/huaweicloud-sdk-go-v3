package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListRepositoryTagsRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType ListRepositoryTagsRequestContentType `json:"Content-Type"`
	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`
	// 起始索引。**注意：offset和limit参数需要配套使用**

	Offset *string `json:"offset,omitempty"`
	// 返回条数。**注意：offset和limit参数需要配套使用*

	Limit *string `json:"limit,omitempty"`
	// 按列排序，可设置为updated_at（按更新时间排序）

	OrderColumn *string `json:"order_column,omitempty"`
	// 排序类型，可设置为desc（降序）、asc（升序）

	OrderType *ListRepositoryTagsRequestOrderType `json:"order_type,omitempty"`
	// 镜像版本名

	Tag *string `json:"tag,omitempty"`
}

func (o ListRepositoryTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryTagsRequest struct{}"
	}

	return strings.Join([]string{"ListRepositoryTagsRequest", string(data)}, " ")
}

type ListRepositoryTagsRequestContentType struct {
	value string
}

type ListRepositoryTagsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListRepositoryTagsRequestContentType
	APPLICATION_JSON             ListRepositoryTagsRequestContentType
}

func GetListRepositoryTagsRequestContentTypeEnum() ListRepositoryTagsRequestContentTypeEnum {
	return ListRepositoryTagsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListRepositoryTagsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListRepositoryTagsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListRepositoryTagsRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryTagsRequestContentType) UnmarshalJSON(b []byte) error {
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

type ListRepositoryTagsRequestOrderType struct {
	value string
}

type ListRepositoryTagsRequestOrderTypeEnum struct {
	DESC ListRepositoryTagsRequestOrderType
	ASC  ListRepositoryTagsRequestOrderType
}

func GetListRepositoryTagsRequestOrderTypeEnum() ListRepositoryTagsRequestOrderTypeEnum {
	return ListRepositoryTagsRequestOrderTypeEnum{
		DESC: ListRepositoryTagsRequestOrderType{
			value: "desc",
		},
		ASC: ListRepositoryTagsRequestOrderType{
			value: "asc",
		},
	}
}

func (c ListRepositoryTagsRequestOrderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepositoryTagsRequestOrderType) UnmarshalJSON(b []byte) error {
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
