package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDryRunConfigResponse Response Object
type UpdateDryRunConfigResponse struct {
	Root           *DryRunConfigDto `json:"root,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateDryRunConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDryRunConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateDryRunConfigResponse", string(data)}, " ")
}
