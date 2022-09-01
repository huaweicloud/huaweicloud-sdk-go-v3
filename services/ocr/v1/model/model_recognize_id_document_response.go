package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RecognizeIdDocumentResponse struct {

	// 调用成功时表示调用结果。  调用失败时此字段为空。
	Result         *interface{} `json:"result,omitempty" xml:"result"`
	HttpStatusCode int          `json:"-"`
}

func (o RecognizeIdDocumentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizeIdDocumentResponse struct{}"
	}

	return strings.Join([]string{"RecognizeIdDocumentResponse", string(data)}, " ")
}
