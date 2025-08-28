package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPluginConfigRequest Request Object
type ListPluginConfigRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 插件供应商  * BOCHA:博查  * AMAP_WEATHER：高德天气
	PluginProvider *ListPluginConfigRequestPluginProvider `json:"plugin_provider,omitempty"`

	// 插件类型  * WEATHER_QUERY：天气查询  * WEB_SEARCH：网络搜索
	PluginType *ListPluginConfigRequestPluginType `json:"plugin_type,omitempty"`
}

func (o ListPluginConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginConfigRequest struct{}"
	}

	return strings.Join([]string{"ListPluginConfigRequest", string(data)}, " ")
}

type ListPluginConfigRequestPluginProvider struct {
	value string
}

type ListPluginConfigRequestPluginProviderEnum struct {
	BOCHA        ListPluginConfigRequestPluginProvider
	AMAP_WEATHER ListPluginConfigRequestPluginProvider
}

func GetListPluginConfigRequestPluginProviderEnum() ListPluginConfigRequestPluginProviderEnum {
	return ListPluginConfigRequestPluginProviderEnum{
		BOCHA: ListPluginConfigRequestPluginProvider{
			value: "BOCHA",
		},
		AMAP_WEATHER: ListPluginConfigRequestPluginProvider{
			value: "AMAP_WEATHER",
		},
	}
}

func (c ListPluginConfigRequestPluginProvider) Value() string {
	return c.value
}

func (c ListPluginConfigRequestPluginProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPluginConfigRequestPluginProvider) UnmarshalJSON(b []byte) error {
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

type ListPluginConfigRequestPluginType struct {
	value string
}

type ListPluginConfigRequestPluginTypeEnum struct {
	WEATHER_QUERY ListPluginConfigRequestPluginType
	WEB_SEARCH    ListPluginConfigRequestPluginType
}

func GetListPluginConfigRequestPluginTypeEnum() ListPluginConfigRequestPluginTypeEnum {
	return ListPluginConfigRequestPluginTypeEnum{
		WEATHER_QUERY: ListPluginConfigRequestPluginType{
			value: "WEATHER_QUERY",
		},
		WEB_SEARCH: ListPluginConfigRequestPluginType{
			value: "WEB_SEARCH",
		},
	}
}

func (c ListPluginConfigRequestPluginType) Value() string {
	return c.value
}

func (c ListPluginConfigRequestPluginType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPluginConfigRequestPluginType) UnmarshalJSON(b []byte) error {
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
