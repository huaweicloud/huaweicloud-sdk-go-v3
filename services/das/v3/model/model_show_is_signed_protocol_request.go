package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIsSignedProtocolRequest Request Object
type ShowIsSignedProtocolRequest struct {
}

func (o ShowIsSignedProtocolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIsSignedProtocolRequest struct{}"
	}

	return strings.Join([]string{"ShowIsSignedProtocolRequest", string(data)}, " ")
}
