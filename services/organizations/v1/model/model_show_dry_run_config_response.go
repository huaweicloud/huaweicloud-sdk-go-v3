package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDryRunConfigResponse Response Object
type ShowDryRunConfigResponse struct {
	Root           *DryRunConfigDto `json:"root,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowDryRunConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDryRunConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowDryRunConfigResponse", string(data)}, " ")
}
