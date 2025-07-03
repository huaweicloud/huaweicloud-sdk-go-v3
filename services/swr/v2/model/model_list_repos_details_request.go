package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListReposDetailsRequest Request Object
type ListReposDetailsRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListReposDetailsRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。注意：和filter最好分开使用，如果同时使用，此过滤参数将失效，以filter为准。
	Namespace *string `json:"namespace,omitempty"`

	// 镜像仓库名称。注意：和filter最好分开使用，如果同时使用，此过滤参数将失效，以filter为准。
	Name *string `json:"name,omitempty"`

	// 镜像仓库分类，可设置为app_server, linux, framework_app, database, lang, other, windows, arm。注意：和filter最好分开使用，如果同时使用，此过滤参数将失效，以filter为准。
	Category *string `json:"category,omitempty"`

	// 返回条数，默认情况下返回100条记录，最多返回1000条记录。注意：offset和limit参数需要配套使用。
	Limit *string `json:"limit,omitempty"`

	// 起始索引。注意：offset和limit参数需要配套使用。
	Offset *string `json:"offset,omitempty"`

	// 按列排序，可设置为updated_at（按更新时间排序）。注意：order_column和order_type参数需要配套使用。
	OrderColumn *string `json:"order_column,omitempty"`

	// 排序类型，可设置为desc（降序）、asc（升序）。注意：order_column和order_type参数需要配套使用。
	OrderType *string `json:"order_type,omitempty"`

	// 注意：如果使用filter至少要传递一个filter参数。应填写 namespace::{namespace}|name::{name}|limit::{limit}|offset::{offset}|order_column::{order_column}|order_type::{order_type},其中{namespace}为组织名称。 {name}为镜像仓库名称，模糊匹配。{category}为镜像仓库分类，可设置为app_server、linux、framework_app、database、lang、other、windows、arm。 {limit}为返回条数,{offset}为起始索引，注意：offset和limit参数需要配套使用。{order_column}为按列排序，可设置为name、updated_time、tag_count,{order_type}为排序类型， 可设置为desc（降序）、asc（升序），注意：order_column和order_type参数需要配套使用。
	Filter *string `json:"filter,omitempty"`
}

func (o ListReposDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReposDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListReposDetailsRequest", string(data)}, " ")
}

type ListReposDetailsRequestContentType struct {
	value string
}

type ListReposDetailsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListReposDetailsRequestContentType
	APPLICATION_JSON             ListReposDetailsRequestContentType
}

func GetListReposDetailsRequestContentTypeEnum() ListReposDetailsRequestContentTypeEnum {
	return ListReposDetailsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListReposDetailsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListReposDetailsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListReposDetailsRequestContentType) Value() string {
	return c.value
}

func (c ListReposDetailsRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListReposDetailsRequestContentType) UnmarshalJSON(b []byte) error {
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
