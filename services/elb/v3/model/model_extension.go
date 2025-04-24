package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Extension struct {

	// ep id
	EpId *string `json:"ep_id,omitempty"`

	// ep service id
	EpServiceId *string `json:"ep_service_id,omitempty"`
}

func (o Extension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Extension struct{}"
	}

	return strings.Join([]string{"Extension", string(data)}, " ")
}
