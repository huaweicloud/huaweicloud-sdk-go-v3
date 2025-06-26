package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadKeystoreResponse Response Object
type UploadKeystoreResponse struct {
	Result *UploadKeystoreResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadKeystoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadKeystoreResponse struct{}"
	}

	return strings.Join([]string{"UploadKeystoreResponse", string(data)}, " ")
}
