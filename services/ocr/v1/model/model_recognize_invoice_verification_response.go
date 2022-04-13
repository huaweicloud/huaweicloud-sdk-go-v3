package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeInvoiceVerificationResponse struct {
	// 调用成功时表示调用结果，详情参见[响应参数](https://support.huaweicloud.com/api-ocr/ocr_03_0134.html#ocr_03_0134__table266mcpsimp)。  调用失败时无此字段。

	Result         *interface{} `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o RecognizeInvoiceVerificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeInvoiceVerificationResponse struct{}"
	}

	return strings.Join([]string{"RecognizeInvoiceVerificationResponse", string(data)}, " ")
}
