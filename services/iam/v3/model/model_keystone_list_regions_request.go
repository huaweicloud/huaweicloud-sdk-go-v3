package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListRegionsRequest Request Object
type KeystoneListRegionsRequest struct {
}

func (o KeystoneListRegionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListRegionsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListRegionsRequest", string(data)}, " ")
}
