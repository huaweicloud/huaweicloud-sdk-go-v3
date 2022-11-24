package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TransferEnterpriseMultiAccountReq struct {

	// 企业子账号的客户ID。您可以调用查询企业子账号列表接口，获取响应参数“id”的返回值。
	CustomerId string `json:"customer_id"`

	// 现金账户总划拨金额。 单位：元。取值大于0且精确到小数点后2位。
	Amount string `json:"amount"`

	// 交易序列号，用于防止重复提交。 如果接口调用方不传此参数的值，则系统自动生成。如果接口调用方传入此参数的值，请采用UUID保证全局唯一。
	TransId *string `json:"trans_id,omitempty"`

	// 账户类型： BALANCE_TYPE_DEBIT：余额账户（默认）BALANCE_TYPE_CREDIT：信用账户
	BalanceType *string `json:"balance_type,omitempty"`

	// 账户到期时间，UTC时间，格式为：2016-03-28T14:45:38Z。 只对信用账户有效，用于限制针对有效期到期时间等于该时间的信用账户余额进行拨款，精确到秒。如果查询信用账户可拨款余额的查询结果没有失效时间，表示永久有效，对于这种账本拨款的情况无需填写。
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o TransferEnterpriseMultiAccountReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferEnterpriseMultiAccountReq struct{}"
	}

	return strings.Join([]string{"TransferEnterpriseMultiAccountReq", string(data)}, " ")
}
