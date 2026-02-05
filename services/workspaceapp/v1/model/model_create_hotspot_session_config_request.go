package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHotspotSessionConfigRequest Request Object
type CreateHotspotSessionConfigRequest struct {
	Body *CreateHotspotSessionConfigReq `json:"body,omitempty"`
}

func (o CreateHotspotSessionConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHotspotSessionConfigRequest struct{}"
	}

	return strings.Join([]string{"CreateHotspotSessionConfigRequest", string(data)}, " ")
}
