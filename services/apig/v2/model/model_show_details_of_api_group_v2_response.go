package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowDetailsOfApiGroupV2Response Response Object
type ShowDetailsOfApiGroupV2Response struct {

	// 编号
	Id string `json:"id"`

	// API分组名称
	Name string `json:"name"`

	// 状态   - 1： 有效
	Status ShowDetailsOfApiGroupV2ResponseStatus `json:"status"`

	// 系统默认分配的子域名
	SlDomain string `json:"sl_domain"`

	// 创建时间
	RegisterTime *sdktime.SdkTime `json:"register_time"`

	// 最近修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time"`

	// 是否已上架云商店： - 1：已上架 - 2：未上架 - 3：审核中  [暂不支持](tag:cmcc,ctc,DT,g42,hk_g42,hk_sbc,hk_tm,hws_eu,hws_ocb,OCB,sbc,tm,hws_hk)
	OnSellStatus int32 `json:"on_sell_status"`

	// 分组上绑定的独立域名列表
	UrlDomains *[]UrlDomain `json:"url_domains,omitempty"`

	// 调试域名是否可以访问，true表示可以访问，false表示禁止访问
	SlDomainAccessEnabled *bool `json:"sl_domain_access_enabled,omitempty"`

	// 系统默认分配的子域名列表
	SlDomains *[]string `json:"sl_domains,omitempty"`

	// 描述
	Remark *string `json:"remark,omitempty"`

	// 流控时长内分组下的API的总访问次数限制，默认不限，请根据服务的负载能力自行设置  暂不支持
	CallLimits *int32 `json:"call_limits,omitempty"`

	// 流控时长  暂不支持
	TimeInterval *int32 `json:"time_interval,omitempty"`

	// 流控的时间单位  暂不支持
	TimeUnit *string `json:"time_unit,omitempty"`

	// 是否为默认分组
	IsDefault *int32 `json:"is_default,omitempty"`

	// 分组版本  - V1：全局分组 - V2：应用级分组  暂不支持，默认为V1
	Version *string `json:"version,omitempty"`

	// 分组归属的集成应用编号。  分组版本V2时必填。  暂不支持
	RomaAppId *string `json:"roma_app_id,omitempty"`

	// 分组归属的集成应用名称  暂不支持
	RomaAppName    *string `json:"roma_app_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailsOfApiGroupV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfApiGroupV2Response struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfApiGroupV2Response", string(data)}, " ")
}

type ShowDetailsOfApiGroupV2ResponseStatus struct {
	value int32
}

type ShowDetailsOfApiGroupV2ResponseStatusEnum struct {
	E_1 ShowDetailsOfApiGroupV2ResponseStatus
}

func GetShowDetailsOfApiGroupV2ResponseStatusEnum() ShowDetailsOfApiGroupV2ResponseStatusEnum {
	return ShowDetailsOfApiGroupV2ResponseStatusEnum{
		E_1: ShowDetailsOfApiGroupV2ResponseStatus{
			value: 1,
		},
	}
}

func (c ShowDetailsOfApiGroupV2ResponseStatus) Value() int32 {
	return c.value
}

func (c ShowDetailsOfApiGroupV2ResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailsOfApiGroupV2ResponseStatus) UnmarshalJSON(b []byte) error {
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
