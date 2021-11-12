package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowBlockchainFlavorsResponse struct {
	EnterpriseSpec *InstanceSpc `json:"enterprise_spec,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowBlockchainFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ShowBlockchainFlavorsResponse", string(data)}, " ")
}
