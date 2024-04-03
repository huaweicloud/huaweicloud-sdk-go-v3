package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeCertificateResponse Response Object
type UnsubscribeCertificateResponse struct {

	// 退订结果。
	UnsubscribeResults *string `json:"unsubscribe_results,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o UnsubscribeCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeCertificateResponse struct{}"
	}

	return strings.Join([]string{"UnsubscribeCertificateResponse", string(data)}, " ")
}
