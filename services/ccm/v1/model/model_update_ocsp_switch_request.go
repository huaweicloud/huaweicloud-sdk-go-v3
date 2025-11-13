package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOcspSwitchRequest Request Object
type UpdateOcspSwitchRequest struct {

	// 所要启用OCSP的CA证书ID。
	CaId string `json:"ca_id"`

	Body *UpdateOcspSwitchRequestBody `json:"body,omitempty"`
}

func (o UpdateOcspSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOcspSwitchRequest struct{}"
	}

	return strings.Join([]string{"UpdateOcspSwitchRequest", string(data)}, " ")
}
