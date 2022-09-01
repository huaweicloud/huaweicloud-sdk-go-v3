package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListReposDetailsRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListReposDetailsRequestContentType `json:"Content-Type" xml:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace *string `json:"namespace,omitempty" xml:"namespace"`

	// 镜像仓库名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 镜像仓库分类，可设置为app_server, linux, framework_app, database, lang, other, windows, arm。
	Category *string `json:"category,omitempty" xml:"category"`

	// 返回条数。注意：offset和limit参数需要配套使用。
	Limit *string `json:"limit,omitempty" xml:"limit"`

	// 起始索引。注意：offset和limit参数需要配套使用。
	Offset *string `json:"offset,omitempty" xml:"offset"`

	// 按列排序，可设置为updated_at（按更新时间排序）。注意：order_column和order_type参数需要配套使用。
	OrderColumn *string `json:"order_column,omitempty" xml:"order_column"`

	// 排序类型，可设置为desc（降序）、asc（升序）。注意：order_column和order_type参数需要配套使用。
	OrderType *string `json:"order_type,omitempty" xml:"order_type"`

	// 注意：至少要传递一个filter参数。其中{namespace}为组织名称，模糊匹配。 {name}为镜像仓库名称，模糊匹配。{category}为镜像仓库分类，可设置为app_server、linux、framework_app、database、lang、other、windows、arm。 {limit}为返回条数,{offset}为起始索引，注意：offset和limit参数需要配套使用。{order_column}为按列排序，可设置为name、updated_time、tag_count,{order_type}为排序类型， 可设置为desc（降序）、asc（升序），注意：order_column和order_type参数需要配套使用。
	Filter *string `json:"filter,omitempty" xml:"filter"`
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
