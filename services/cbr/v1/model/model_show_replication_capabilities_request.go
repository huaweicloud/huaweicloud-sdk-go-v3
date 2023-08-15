package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReplicationCapabilitiesRequest Request Object
type ShowReplicationCapabilitiesRequest struct {
}

func (o ShowReplicationCapabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationCapabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ShowReplicationCapabilitiesRequest", string(data)}, " ")
}
