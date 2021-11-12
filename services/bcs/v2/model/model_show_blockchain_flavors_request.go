package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBlockchainFlavorsRequest struct {
}

func (o ShowBlockchainFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainFlavorsRequest", string(data)}, " ")
}
