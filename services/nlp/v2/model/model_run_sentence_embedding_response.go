package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RunSentenceEmbeddingResponse struct {

	// 句向量结果列表，按输入句子顺序返回句向量。调用失败时无此字段。
	Vectors *[][]float32 `json:"vectors,omitempty"`

	// 调用失败时的错误码，具体请参见错误码。调用成功时无此字段。
	ErrorCode *string `json:"error_code,omitempty"`

	// 调用失败时的错误信息。调用成功时无此字段。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunSentenceEmbeddingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunSentenceEmbeddingResponse struct{}"
	}

	return strings.Join([]string{"RunSentenceEmbeddingResponse", string(data)}, " ")
}
