package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGlobalEipRequest Request Object
type UpdateGlobalEipRequest struct {
	GlobalEipId string `json:"global_eip_id"`

	Body *UpdateGlobalEipRequestBody `json:"body,omitempty"`
}

func (o UpdateGlobalEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGlobalEipRequest struct{}"
	}

	return strings.Join([]string{"UpdateGlobalEipRequest", string(data)}, " ")
}
