package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetDocumentAtomicInfoResponse Response Object
type GetDocumentAtomicInfoResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 返回数据。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o GetDocumentAtomicInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetDocumentAtomicInfoResponse struct{}"
	}

	return strings.Join([]string{"GetDocumentAtomicInfoResponse", string(data)}, " ")
}
