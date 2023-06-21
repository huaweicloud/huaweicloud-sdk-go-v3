package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 生成短链的响应参数对象。
type ListResolveTaskResultParam struct {

	// 创建解析任务时填写用户唯一标识，手机号码或者任何的唯一标识，唯一标识不超过64个字符。 发送智能信息时则必须填客户的手机号码。样例为：130****0001。
	CustFlag *string `json:"cust_flag,omitempty"`

	// 租户ID。
	CustId *string `json:"cust_id,omitempty"`

	// 动态参数。
	DyncParams map[string]string `json:"dync_params,omitempty"`

	// 自定义跳转地址。 > - 未填时，终端用户点击访问短信原文中的短链，跳转智能信息H5页 > - 已填时，终端用户点击访问短信原文中的短链，跳转客户填写的链接落地页，填写时必须为http或https作为前缀
	CustomUrl *string `json:"custom_url,omitempty"`

	// 完整的短链连接地址，通过自己的短信渠道发送时，需要把该短链添加到短信模板中，并确保发送短信时的签名与创建短链时的签名保持一致。样例：km2g.cn/PDiWqc。
	AimUrl *string `json:"aim_url,omitempty"`

	// 智能信息编码。样例：PDiWqc。
	AimCode *string `json:"aim_code,omitempty"`

	// 自定义扩展参数。  >预留字段。
	ExtData *string `json:"ext_data,omitempty"`

	// 短链申请结果返回码。 - 0：成功 - 非0：失败，具体请参见错误码
	ResultCode *string `json:"result_code,omitempty"`

	// 短链申请结果错误描述。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 短链生成时间。样例为：2019-10-12T07:20:50。
	GenerateDate *string `json:"generate_date,omitempty"`

	// 短链到期时间。样例为：2019-10-12T07:20:50。
	ExpireDate *string `json:"expire_date,omitempty"`

	// 解析时间。样例为：2019-10-12T07:20:50。
	ResolvedDate *string `json:"resolved_date,omitempty"`

	// 短链实际解析次数。  > 预留字段。
	ResolvedTimes *int32 `json:"resolved_times,omitempty"`
}

func (o ListResolveTaskResultParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResolveTaskResultParam struct{}"
	}

	return strings.Join([]string{"ListResolveTaskResultParam", string(data)}, " ")
}
