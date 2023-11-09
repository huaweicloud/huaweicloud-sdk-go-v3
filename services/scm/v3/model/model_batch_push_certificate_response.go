package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchPushCertificateResponse Response Object
type BatchPushCertificateResponse struct {

	// 部署结果。
	Results        *[]BatchPushCertificateResponseBodyResults `json:"results,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o BatchPushCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPushCertificateResponse struct{}"
	}

	return strings.Join([]string{"BatchPushCertificateResponse", string(data)}, " ")
}
