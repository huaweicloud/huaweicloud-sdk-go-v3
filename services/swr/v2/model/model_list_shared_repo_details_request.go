package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSharedRepoDetailsRequest Request Object
type ListSharedRepoDetailsRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType ListSharedRepoDetailsRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace *string `json:"namespace,omitempty"`

	// 镜像仓库名称
	Name *string `json:"name,omitempty"`

	// self: 我共享的镜像。thirdparty: 他人共享给我的镜像
	SharedBy ListSharedRepoDetailsRequestSharedBy `json:"shared_by"`

	// 返回条数,默认返回100条记录，最多返回1000条。
	Limit *int64 `json:"limit,omitempty"`

	// 分页查询时下一次查询的起始地址。
	Marker *int64 `json:"marker,omitempty"`

	// 查询他人共享给我的镜像是否已过期 false：共享已过期。true：正常
	Status *bool `json:"status,omitempty"`
}

func (o ListSharedRepoDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharedRepoDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListSharedRepoDetailsRequest", string(data)}, " ")
}

type ListSharedRepoDetailsRequestContentType struct {
	value string
}

type ListSharedRepoDetailsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListSharedRepoDetailsRequestContentType
	APPLICATION_JSON             ListSharedRepoDetailsRequestContentType
}

func GetListSharedRepoDetailsRequestContentTypeEnum() ListSharedRepoDetailsRequestContentTypeEnum {
	return ListSharedRepoDetailsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListSharedRepoDetailsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListSharedRepoDetailsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListSharedRepoDetailsRequestContentType) Value() string {
	return c.value
}

func (c ListSharedRepoDetailsRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSharedRepoDetailsRequestContentType) UnmarshalJSON(b []byte) error {
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

type ListSharedRepoDetailsRequestSharedBy struct {
	value string
}

type ListSharedRepoDetailsRequestSharedByEnum struct {
	SELF       ListSharedRepoDetailsRequestSharedBy
	THIRDPARTY ListSharedRepoDetailsRequestSharedBy
}

func GetListSharedRepoDetailsRequestSharedByEnum() ListSharedRepoDetailsRequestSharedByEnum {
	return ListSharedRepoDetailsRequestSharedByEnum{
		SELF: ListSharedRepoDetailsRequestSharedBy{
			value: "self",
		},
		THIRDPARTY: ListSharedRepoDetailsRequestSharedBy{
			value: "thirdparty",
		},
	}
}

func (c ListSharedRepoDetailsRequestSharedBy) Value() string {
	return c.value
}

func (c ListSharedRepoDetailsRequestSharedBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSharedRepoDetailsRequestSharedBy) UnmarshalJSON(b []byte) error {
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
