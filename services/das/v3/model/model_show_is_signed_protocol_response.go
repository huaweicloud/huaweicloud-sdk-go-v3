package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIsSignedProtocolResponse Response Object
type ShowIsSignedProtocolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ShowIsSignedProtocolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIsSignedProtocolResponse struct{}"
	}

	return strings.Join([]string{"ShowIsSignedProtocolResponse", string(data)}, " ")
}
