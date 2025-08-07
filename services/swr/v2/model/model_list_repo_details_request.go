package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListRepoDetailsRequest Request Object
type ListRepoDetailsRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListRepoDetailsRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace *string `json:"namespace,omitempty"`

	// 镜像仓库名称。
	Name *string `json:"name,omitempty"`

	// 镜像仓库分类，可设置为app_server, linux, framework_app, database, lang, other, windows, arm。
	Category *string `json:"category,omitempty"`

	// 返回条数，默认返回100条记录，最多返回1000条记录。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询下一次查询起始标记，接口的返回值nextMarker为下一次查询的起始标记。
	Marker *int32 `json:"marker,omitempty"`

	// 是否公开私有，true为公开，false为私有。
	IsPublic *bool `json:"is_public,omitempty"`
}

func (o ListRepoDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepoDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListRepoDetailsRequest", string(data)}, " ")
}

type ListRepoDetailsRequestContentType struct {
	value string
}

type ListRepoDetailsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListRepoDetailsRequestContentType
	APPLICATION_JSON             ListRepoDetailsRequestContentType
}

func GetListRepoDetailsRequestContentTypeEnum() ListRepoDetailsRequestContentTypeEnum {
	return ListRepoDetailsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListRepoDetailsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListRepoDetailsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListRepoDetailsRequestContentType) Value() string {
	return c.value
}

func (c ListRepoDetailsRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListRepoDetailsRequestContentType) UnmarshalJSON(b []byte) error {
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
