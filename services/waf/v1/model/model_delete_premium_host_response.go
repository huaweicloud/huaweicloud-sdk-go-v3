package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeletePremiumHostResponse Response Object
type DeletePremiumHostResponse struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名
	Hostname *string `json:"hostname,omitempty"`

	// 扩展字段，用于保存防护域名的一些配置信息。
	Extend map[string]string `json:"extend,omitempty"`

	// 华为云区域ID，控制台创建的域名会携带此参数，api调用创建的域名此参数为空，可以通过华为云上地区和终端节点文档查询区域ID对应的中文名称
	Region *string `json:"region,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 域名描述
	Description *string `json:"description,omitempty"`

	// 防护域名初始绑定的防护策略ID,可以通过策略名称调用查询防护策略列表（ListPolicy）接口查询到对应的策略id
	Policyid *string `json:"policyid,omitempty"`

	// **参数解释：** 域名防护状态标识，用于指定域名在WAF中的防护运行状态 **约束限制：** 不涉及 **取值范围：**  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测  **默认取值：** 不涉及
	ProtectStatus *DeletePremiumHostResponseProtectStatus `json:"protect_status,omitempty"`

	// **参数解释：** 域名接入状态 **约束限制：** 不涉及 **取值范围：**  - 0: 未接入  - 1: 已接入  **默认取值：** 不涉及
	AccessStatus *DeletePremiumHostResponseAccessStatus `json:"access_status,omitempty"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty"`

	// 域名id，和id的值是一样的，属于冗余字段
	HostId         *string `json:"host_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeletePremiumHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePremiumHostResponse struct{}"
	}

	return strings.Join([]string{"DeletePremiumHostResponse", string(data)}, " ")
}

type DeletePremiumHostResponseProtectStatus struct {
	value int32
}

type DeletePremiumHostResponseProtectStatusEnum struct {
	E_MINUS_1 DeletePremiumHostResponseProtectStatus
	E_0       DeletePremiumHostResponseProtectStatus
	E_1       DeletePremiumHostResponseProtectStatus
}

func GetDeletePremiumHostResponseProtectStatusEnum() DeletePremiumHostResponseProtectStatusEnum {
	return DeletePremiumHostResponseProtectStatusEnum{
		E_MINUS_1: DeletePremiumHostResponseProtectStatus{
			value: -1,
		}, E_0: DeletePremiumHostResponseProtectStatus{
			value: 0,
		}, E_1: DeletePremiumHostResponseProtectStatus{
			value: 1,
		},
	}
}

func (c DeletePremiumHostResponseProtectStatus) Value() int32 {
	return c.value
}

func (c DeletePremiumHostResponseProtectStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeletePremiumHostResponseProtectStatus) UnmarshalJSON(b []byte) error {
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

type DeletePremiumHostResponseAccessStatus struct {
	value int32
}

type DeletePremiumHostResponseAccessStatusEnum struct {
	E_0 DeletePremiumHostResponseAccessStatus
	E_1 DeletePremiumHostResponseAccessStatus
}

func GetDeletePremiumHostResponseAccessStatusEnum() DeletePremiumHostResponseAccessStatusEnum {
	return DeletePremiumHostResponseAccessStatusEnum{
		E_0: DeletePremiumHostResponseAccessStatus{
			value: 0,
		}, E_1: DeletePremiumHostResponseAccessStatus{
			value: 1,
		},
	}
}

func (c DeletePremiumHostResponseAccessStatus) Value() int32 {
	return c.value
}

func (c DeletePremiumHostResponseAccessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeletePremiumHostResponseAccessStatus) UnmarshalJSON(b []byte) error {
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
