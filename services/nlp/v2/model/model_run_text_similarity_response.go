package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunTextSimilarityResponse struct {
	// 相似度得分，范围在0~1，默认小数点后保留8位。调用失败时无此字段。

	Similarity *float32 `json:"similarity,omitempty"`
	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。

	ErrorCode *string `json:"error_code,omitempty"`
	// 调用失败时的错误信息。调用成功时无此字段。

	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunTextSimilarityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunTextSimilarityResponse struct{}"
	}

	return strings.Join([]string{"RunTextSimilarityResponse", string(data)}, " ")
}
