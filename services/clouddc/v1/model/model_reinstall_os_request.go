package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReinstallOsRequest Request Object
type ReinstallOsRequest struct {
	Body *ReinstallOsOptions `json:"body,omitempty"`
}

func (o ReinstallOsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallOsRequest struct{}"
	}

	return strings.Join([]string{"ReinstallOsRequest", string(data)}, " ")
}
