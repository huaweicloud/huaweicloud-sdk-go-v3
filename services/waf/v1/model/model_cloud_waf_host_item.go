package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CloudWafHostItem struct {

	// 域名id
	Id *string `json:"id,omitempty"`

	// 域名id
	Hostid *string `json:"hostid,omitempty"`

	// 华为云区域ID，控制台创建的域名会携带此参数，api调用创建的域名此参数为空，可以通过华为云上地区和终端节点文档查询区域ID对应的中文名称
	Region *string `json:"region,omitempty"`

	// 域名描述信息，可选参数。
	Description *string `json:"description,omitempty"`

	// 防护域名的源站服务器配置信息
	Server *[]CloudWafServer `json:"server,omitempty"`

	// WAF部署模式，默认是1，目前仅支持反代模式，冗余参数
	Type *int32 `json:"type,omitempty"`

	// 防护域名是否使用代理   - false：不使用代理   - true：使用代理
	Proxy *bool `json:"proxy,omitempty"`

	// 创建的云模式防护域名
	Hostname *string `json:"hostname,omitempty"`

	// cname前缀
	AccessCode *string `json:"access_code,omitempty"`

	// 防护策略id
	Policyid *string `json:"policyid,omitempty"`

	// 创建防护域名的时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// **参数解释：** 域名防护状态标识，用于指定域名在WAF中的防护运行状态 **约束限制：** 不涉及 **取值范围：**  - -1：bypass，该域名的请求直接到达其后端服务器，不再经过WAF  - 0：暂停防护，WAF只转发该域名的请求，不做攻击检测  - 1：开启防护，WAF根据您配置的策略进行攻击检测  **默认取值：** 不涉及
	ProtectStatus *int32 `json:"protect_status,omitempty"`

	// **参数解释：** 域名接入状态 **约束限制：** 不涉及 **取值范围：**  - 0: 未接入  - 1: 已接入  **默认取值：** 不涉及
	AccessStatus *int32 `json:"access_status,omitempty"`

	// 是否使用独享ip   - true：使用独享ip   - false：不使用独享ip
	ExclusiveIp *bool `json:"exclusive_ip,omitempty"`

	// **参数解释：** 套餐付费模式标识，用于指定套餐的计费方式 **约束限制：** 不涉及 **取值范围：**  - prePaid:包周期模式  - postPaid:按需模式  **默认取值：** prePaid
	PaidType *CloudWafHostItemPaidType `json:"paid_type,omitempty"`

	// 网站名称，对应WAF控制台域名详情中的网站名称
	WebTag *string `json:"web_tag,omitempty"`

	Flag *Flag `json:"flag,omitempty"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o CloudWafHostItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudWafHostItem struct{}"
	}

	return strings.Join([]string{"CloudWafHostItem", string(data)}, " ")
}

type CloudWafHostItemPaidType struct {
	value string
}

type CloudWafHostItemPaidTypeEnum struct {
	PRE_PAID  CloudWafHostItemPaidType
	POST_PAID CloudWafHostItemPaidType
}

func GetCloudWafHostItemPaidTypeEnum() CloudWafHostItemPaidTypeEnum {
	return CloudWafHostItemPaidTypeEnum{
		PRE_PAID: CloudWafHostItemPaidType{
			value: "prePaid",
		},
		POST_PAID: CloudWafHostItemPaidType{
			value: "postPaid",
		},
	}
}

func (c CloudWafHostItemPaidType) Value() string {
	return c.value
}

func (c CloudWafHostItemPaidType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CloudWafHostItemPaidType) UnmarshalJSON(b []byte) error {
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
