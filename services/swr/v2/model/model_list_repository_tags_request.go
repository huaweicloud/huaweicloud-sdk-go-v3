package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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

	// 返回条数。注意：offset和limit参数需要配套使用。
	Limit *string `json:"limit,omitempty"`

	// 起始索引。注意：offset和limit参数需要配套使用。
	Offset *string `json:"offset,omitempty"`

	// 按列排序，可设置为updated_at（按更新时间排序）。注意：order_column和order_type参数需要配套使用。
	OrderColumn *string `json:"order_column,omitempty"`

	// 排序类型，可设置为desc（降序）、asc（升序）。注意：order_column和order_type参数需要配套使用。
	OrderType *string `json:"order_type,omitempty"`

	// 镜像版本名。
	Tag *string `json:"tag,omitempty"`

	// 应填写 offset::{offset}|limit::{limit}|order_column::{order_column}|order_type::{order_type}|tag::{tag} ,其中{limit}为返回条数,{offset}为起始索引,注意：offset和limit参数需要配套使用。 {order_column}为按列排序，可设置为updated_at（按更新时间排序）,{order_type}为排序类型，可设置为desc（降序）、asc（升序），{tag}为镜像版本名。
	Filter *string `json:"filter,omitempty"`
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

func (c ListRepositoryTagsRequestContentType) Value() string {
	return c.value
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
