package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizeSmartDocumentRecognizerResponse Response Object
type RecognizeSmartDocumentRecognizerResponse struct {

	// 调用成功时返回的结果列表，按页面顺序返回，列表第一项为第一页识别结果，依次类推。 调用失败时无此字段。
	Result *[]SmartDocumentRecognizerResult `json:"result,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RecognizeSmartDocumentRecognizerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeSmartDocumentRecognizerResponse struct{}"
	}

	return strings.Join([]string{"RecognizeSmartDocumentRecognizerResponse", string(data)}, " ")
}
