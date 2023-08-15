package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListAuthProjectsRequest Request Object
type KeystoneListAuthProjectsRequest struct {
}

func (o KeystoneListAuthProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListAuthProjectsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListAuthProjectsRequest", string(data)}, " ")
}
