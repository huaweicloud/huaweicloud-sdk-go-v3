package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeInvoiceVerificationResponse Response Object
type RecognizeInvoiceVerificationResponse struct {

	// 调用成功时表示调用结果，详情参见[响应参数](https://support.huaweicloud.com/api-ocr/ocr_03_0134.html#ocr_03_0134__table266mcpsimp)。  调用失败时无此字段。 依据发票类型不同，返回参数不同。 - 增值税发票   含增值税专用发票、增值税普通发票、增值税普通发票（卷式）、增值税电子专用发票、增值税电子普通发票、增值税电子普通发票（通行费）、区块链电子发票。 - 机动车销售统一发票 - 二手车销售统一发票
	Result *interface{} `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeInvoiceVerificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeInvoiceVerificationResponse struct{}"
	}

	return strings.Join([]string{"RecognizeInvoiceVerificationResponse", string(data)}, " ")
}
