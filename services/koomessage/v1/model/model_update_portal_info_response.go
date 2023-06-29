package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePortalInfoResponse Response Object
type UpdatePortalInfoResponse struct {
	Data           *UpdatePortalInfoResponseModel `json:"data,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o UpdatePortalInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortalInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdatePortalInfoResponse", string(data)}, " ")
}
