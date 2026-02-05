package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListApisV2Request Request Object
type ListApisV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty"`

	// 每页显示的条目数量，条目数量小于等于0时，自动转换为20，条目数量大于500时，自动转换为500
	Limit *int32 `json:"limit,omitempty"`

	// API编号
	Id *string `json:"id,omitempty"`

	// API名称
	Name *string `json:"name,omitempty"`

	// API分组编号
	GroupId *string `json:"group_id,omitempty"`

	// 请求协议
	ReqProtocol *string `json:"req_protocol,omitempty"`

	// 请求方法
	ReqMethod *string `json:"req_method,omitempty"`

	// 请求路径
	ReqUri *string `json:"req_uri,omitempty"`

	// 授权类型。 - NONE：无认证 - APP：APP认证 - IAM：IAM认证 - AUTHORIZER：自定义认证
	AuthType *ListApisV2RequestAuthType `json:"auth_type,omitempty"`

	// 发布的环境编号
	EnvId *string `json:"env_id,omitempty"`

	// API类型。 - 1：公有API - 2：私有API
	Type *ListApisV2RequestType `json:"type,omitempty"`

	// 指定需要精确匹配查找的参数名称，目前仅支持name、req_uri
	PreciseSearch *string `json:"precise_search,omitempty"`

	// 负载通道名称
	VpcChannelName *string `json:"vpc_channel_name,omitempty"`

	// 指定API详情中需要包含的额外返回结果，多个参数之间使用“,”隔开，当brief和其他include参数共同使用时，brief不生效。 目前仅支持brief，include_group，include_group_backend。 brief：默认值，不包含额外信息。 include_group：返回结果中包含api_group_info。 include_group_backend：返回结果中包含backend_api。
	ReturnDataMode *string `json:"return_data_mode,omitempty"`

	// API标签，该参数可指定多个，多个不同的参数值为或关系；不指定或为空时，表示不筛选标签；指定为#no_tags#时，表示筛选无标签API。
	Tags *string `json:"tags,omitempty"`
}

func (o ListApisV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApisV2Request struct{}"
	}

	return strings.Join([]string{"ListApisV2Request", string(data)}, " ")
}

type ListApisV2RequestAuthType struct {
	value string
}

type ListApisV2RequestAuthTypeEnum struct {
	NONE       ListApisV2RequestAuthType
	APP        ListApisV2RequestAuthType
	IAM        ListApisV2RequestAuthType
	AUTHORIZER ListApisV2RequestAuthType
}

func GetListApisV2RequestAuthTypeEnum() ListApisV2RequestAuthTypeEnum {
	return ListApisV2RequestAuthTypeEnum{
		NONE: ListApisV2RequestAuthType{
			value: "NONE",
		},
		APP: ListApisV2RequestAuthType{
			value: "APP",
		},
		IAM: ListApisV2RequestAuthType{
			value: "IAM",
		},
		AUTHORIZER: ListApisV2RequestAuthType{
			value: "AUTHORIZER",
		},
	}
}

func (c ListApisV2RequestAuthType) Value() string {
	return c.value
}

func (c ListApisV2RequestAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApisV2RequestAuthType) UnmarshalJSON(b []byte) error {
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

type ListApisV2RequestType struct {
	value int32
}

type ListApisV2RequestTypeEnum struct {
	E_1 ListApisV2RequestType
	E_2 ListApisV2RequestType
}

func GetListApisV2RequestTypeEnum() ListApisV2RequestTypeEnum {
	return ListApisV2RequestTypeEnum{
		E_1: ListApisV2RequestType{
			value: 1,
		}, E_2: ListApisV2RequestType{
			value: 2,
		},
	}
}

func (c ListApisV2RequestType) Value() int32 {
	return c.value
}

func (c ListApisV2RequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListApisV2RequestType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
