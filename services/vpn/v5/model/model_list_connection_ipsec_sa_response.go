package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionIpsecSaResponse Response Object
type ListConnectionIpsecSaResponse struct {
	SaInfos *[]SaInfo `json:"sa_infos,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ListConnectionIpsecSaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionIpsecSaResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionIpsecSaResponse", string(data)}, " ")
}
