package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePublicipRequest struct {
	// 弹性公网IP唯一标识

	PublicipId string `json:"publicip_id"`

	Body *UpdatePublicipsRequestBody `json:"body,omitempty"`
}

func (o UpdatePublicipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicipRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicipRequest", string(data)}, " ")
}
