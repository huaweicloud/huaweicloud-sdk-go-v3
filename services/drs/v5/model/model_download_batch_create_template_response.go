package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadBatchCreateTemplateResponse Response Object
type DownloadBatchCreateTemplateResponse struct {

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述。
	ErrorMsg       *string `json:"error_msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DownloadBatchCreateTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBatchCreateTemplateResponse struct{}"
	}

	return strings.Join([]string{"DownloadBatchCreateTemplateResponse", string(data)}, " ")
}
