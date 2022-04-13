package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Request Object
type ListSharedReposDetailsRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType ListSharedReposDetailsRequestContentType `json:"Content-Type"`
	// 应填写 center::{center}|limit::{limit}|offset::{offset}|order_column::{order_column}|order_type::{order_type} ,其中{limit}为返回条数,{offset}为起始索引, {order_column}为按列排序，可设置为name、updated_time、tag_count,{order_type}为排序类型，可设置为desc（降序）、asc（升序）

	Filter *string `json:"filter,omitempty"`
}

func (o ListSharedReposDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharedReposDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSharedReposDetailsRequest", string(data)}, " ")
}

type ListSharedReposDetailsRequestContentType struct {
	value string
}

type ListSharedReposDetailsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListSharedReposDetailsRequestContentType
	APPLICATION_JSON             ListSharedReposDetailsRequestContentType
}

func GetListSharedReposDetailsRequestContentTypeEnum() ListSharedReposDetailsRequestContentTypeEnum {
	return ListSharedReposDetailsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListSharedReposDetailsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListSharedReposDetailsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListSharedReposDetailsRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSharedReposDetailsRequestContentType) UnmarshalJSON(b []byte) error {
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
