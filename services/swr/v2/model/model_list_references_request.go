package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListReferencesRequest Request Object
type ListReferencesRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListReferencesRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace"`

	// 镜像仓库名称
	Repository string `json:"repository"`

	// 签名镜像的版本号
	Tag string `json:"tag"`

	// 返回条数。如果不传该参数默认返回10条记录，最大支持10条记录
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询时的起始标记，接口的返回值next_marker为下一次查询的起始标记
	Marker *string `json:"marker,omitempty"`
}

func (o ListReferencesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListReferencesRequest struct{}"
	}

	return strings.Join([]string{"ListReferencesRequest", string(data)}, " ")
}

type ListReferencesRequestContentType struct {
	value string
}

type ListReferencesRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListReferencesRequestContentType
	APPLICATION_JSON             ListReferencesRequestContentType
}

func GetListReferencesRequestContentTypeEnum() ListReferencesRequestContentTypeEnum {
	return ListReferencesRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListReferencesRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListReferencesRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListReferencesRequestContentType) Value() string {
	return c.value
}

func (c ListReferencesRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListReferencesRequestContentType) UnmarshalJSON(b []byte) error {
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
