/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"strings"
)

// Response Object
type ShowDetailsOfAppV2Response struct {
	// APP的创建者 - USER：用户自行创建 - MARKET：云市场分配
	Creator ShowDetailsOfAppV2ResponseCreator `json:"creator,omitempty"`
	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
	// APP的key
	AppKey *string `json:"app_key,omitempty"`
	// 名称
	Name *string `json:"name,omitempty"`
	// 描述
	Remark *string `json:"remark,omitempty"`
	// 编号
	Id *string `json:"id,omitempty"`
	// 密钥
	AppSecret *string `json:"app_secret,omitempty"`
	// 注册时间
	RegisterTime *sdktime.SdkTime `json:"register_time,omitempty"`
	// 状态
	Status *int32 `json:"status,omitempty"`
	// APP的类型  默认为apig，暂不支持其他类型
	AppType ShowDetailsOfAppV2ResponseAppType `json:"app_type,omitempty"`
}

func (o ShowDetailsOfAppV2Response) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowDetailsOfAppV2Response", string(data)}, " ")
}

type ShowDetailsOfAppV2ResponseCreator struct {
	value string
}

type ShowDetailsOfAppV2ResponseCreatorEnum struct {
	USER   ShowDetailsOfAppV2ResponseCreator
	MARKET ShowDetailsOfAppV2ResponseCreator
}

func GetShowDetailsOfAppV2ResponseCreatorEnum() ShowDetailsOfAppV2ResponseCreatorEnum {
	return ShowDetailsOfAppV2ResponseCreatorEnum{
		USER: ShowDetailsOfAppV2ResponseCreator{
			value: "USER",
		},
		MARKET: ShowDetailsOfAppV2ResponseCreator{
			value: "MARKET",
		},
	}
}

func (c ShowDetailsOfAppV2ResponseCreator) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowDetailsOfAppV2ResponseCreator) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}

type ShowDetailsOfAppV2ResponseAppType struct {
	value string
}

type ShowDetailsOfAppV2ResponseAppTypeEnum struct {
	APIG ShowDetailsOfAppV2ResponseAppType
	ROMA ShowDetailsOfAppV2ResponseAppType
}

func GetShowDetailsOfAppV2ResponseAppTypeEnum() ShowDetailsOfAppV2ResponseAppTypeEnum {
	return ShowDetailsOfAppV2ResponseAppTypeEnum{
		APIG: ShowDetailsOfAppV2ResponseAppType{
			value: "apig",
		},
		ROMA: ShowDetailsOfAppV2ResponseAppType{
			value: "roma",
		},
	}
}

func (c ShowDetailsOfAppV2ResponseAppType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowDetailsOfAppV2ResponseAppType) UnmarshalJSON(b []byte) error {
	c.value = string(strings.Trim(string(b[:]), "\""))
	return nil
}
