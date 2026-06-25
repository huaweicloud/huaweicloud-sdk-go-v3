package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpsConfigRequest Request Object
type UpdateHttpsConfigRequest struct {
	Body *ConfigCdnHttpsReq `json:"body,omitempty"`
}

func (o UpdateHttpsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpsConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpsConfigRequest", string(data)}, " ")
}
