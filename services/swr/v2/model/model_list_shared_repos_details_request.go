package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSharedReposDetailsRequest Request Object
type ListSharedReposDetailsRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListSharedReposDetailsRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace *string `json:"namespace,omitempty"`

	// 镜像仓库名称
	Name *string `json:"name,omitempty"`

	// self: 我共享的镜像。thirdparty: 他人共享给我的镜像
	Center string `json:"center"`

	// 返回条数。注意：offset和limit参数需要配套使用。
	Limit *string `json:"limit,omitempty"`

	// 起始索引。注意：offset和limit参数需要配套使用。
	Offset *string `json:"offset,omitempty"`

	// 按列排序，可设置为updated_at（按更新时间排序）。注意：order_column和order_type参数需要配套使用。
	OrderColumn *string `json:"order_column,omitempty"`

	// 排序类型，可设置为desc（降序）、asc（升序）。注意：order_column和order_type参数需要配套使用。
	OrderType *string `json:"order_type,omitempty"`

	// 应填写 center::{center}|name::{name}|limit::{limit}|offset::{offset}|namespace::{namespace}|order_column::{order_column}|order_type::{order_type} ,其中 {center}可选为self: 我共享的镜像。thirdparty: 他人共享给我的镜像，namespace为组织名称，name为镜像仓库名称， {limit}为返回条数,{offset}为起始索引,{order_column}为按列排序，可设置为name、updated_time、tag_count,{order_type}为排序类型，可设置为desc（降序）、asc（升序）
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

func (c ListSharedReposDetailsRequestContentType) Value() string {
	return c.value
}

func (c ListSharedReposDetailsRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSharedReposDetailsRequestContentType) UnmarshalJSON(b []byte) error {
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
