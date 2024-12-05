package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListFederationProjectsRequest Request Object
type KeystoneListFederationProjectsRequest struct {
}

func (o KeystoneListFederationProjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListFederationProjectsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListFederationProjectsRequest", string(data)}, " ")
}
