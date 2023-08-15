package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListVersionsRequest Request Object
type KeystoneListVersionsRequest struct {
}

func (o KeystoneListVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListVersionsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListVersionsRequest", string(data)}, " ")
}
