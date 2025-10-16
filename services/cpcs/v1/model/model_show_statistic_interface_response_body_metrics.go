package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowStatisticInterfaceResponseBodyMetrics struct {

	// 杂凑
	PkiDigestBusiStatistic *float64 `json:"pkiDigestBusiStatistic,omitempty"`

	// 解密
	PkiDecryptBusiStatistic *float64 `json:"pkiDecryptBusiStatistic,omitempty"`

	// 生成密钥
	PkiGenerateKeyBusiStatistic *float64 `json:"pkiGenerateKeyBusiStatistic,omitempty"`

	// 验章次数
	PkiVerifyBusiStatistic *float64 `json:"pkiVerifyBusiStatistic,omitempty"`

	// 加密
	PkiEncryptBusiStatistic *float64 `json:"pkiEncryptBusiStatistic,omitempty"`

	// 验证次数
	TsaVerifyBusiStatistic *float64 `json:"tsaVerifyBusiStatistic,omitempty"`

	// 解析次数
	TsaParseBusiStatistic *float64 `json:"tsaParseBusiStatistic,omitempty"`

	// svs验签
	SvsVerifyBusiStatistic *float64 `json:"svsVerifyBusiStatistic,omitempty"`

	// 生成随机
	PkiRandomBusiStatistic *float64 `json:"pkiRandomBusiStatistic,omitempty"`

	// 统计时间，UNIX时间戳，单位毫秒
	Timestamp *int64 `json:"timestamp,omitempty"`

	// svs获取证书等相关证书操作
	SvsCertBusiStatistic *float64 `json:"svsCertBusiStatistic,omitempty"`

	// svs编码
	SvsEncodeBusiStatistic *float64 `json:"svsEncodeBusiStatistic,omitempty"`

	// 申请次数
	TsaApplyBusiStatistic *float64 `json:"tsaApplyBusiStatistic,omitempty"`

	// svs签名
	SvsSignBusiStatistic *float64 `json:"svsSignBusiStatistic,omitempty"`

	// svs解密
	SvsDecryptBusiStatistic *float64 `json:"svsDecryptBusiStatistic,omitempty"`

	// 调用次数
	KmsBusiStatistic *float64 `json:"kmsBusiStatistic,omitempty"`

	// 验章次数
	SealVerifyBusiStatistic *float64 `json:"sealVerifyBusiStatistic,omitempty"`

	// 签名
	PkiSignBusiStatistic *float64 `json:"pkiSignBusiStatistic,omitempty"`

	// 口令验证
	SecauthBusiStatistic *float64 `json:"secauthBusiStatistic,omitempty"`

	// 签章次数
	SealSignBusiStatistic *float64 `json:"sealSignBusiStatistic,omitempty"`

	// 签名次数
	SplitBusiStatistic *float64 `json:"splitBusiStatistic,omitempty"`

	// svs生成随机
	SvsRandomBusiStatistic *float64 `json:"svsRandomBusiStatistic,omitempty"`

	// svs加密
	SvsEncryptBusiStatistic *float64 `json:"svsEncryptBusiStatistic,omitempty"`

	// 解密次数
	SmsDecBusiStatistic *float64 `json:"smsDecBusiStatistic,omitempty"`

	// svs杂凑
	SvsDigestBusiStatistic *float64 `json:"svsDigestBusiStatistic,omitempty"`

	// svs解码
	SvsDecodeBusiStatistic *float64 `json:"svsDecodeBusiStatistic,omitempty"`
}

func (o ShowStatisticInterfaceResponseBodyMetrics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticInterfaceResponseBodyMetrics struct{}"
	}

	return strings.Join([]string{"ShowStatisticInterfaceResponseBodyMetrics", string(data)}, " ")
}
