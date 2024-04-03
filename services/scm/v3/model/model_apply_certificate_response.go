package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyCertificateResponse Response Object
type ApplyCertificateResponse struct {

	// 请求结果。
	RequestInfo    *string `json:"request_info,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ApplyCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyCertificateResponse struct{}"
	}

	return strings.Join([]string{"ApplyCertificateResponse", string(data)}, " ")
}
