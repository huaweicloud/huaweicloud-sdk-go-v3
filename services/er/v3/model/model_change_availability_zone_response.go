package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAvailabilityZoneResponse Response Object
type ChangeAvailabilityZoneResponse struct {
	Instance *EnterpriseRouter `json:"instance,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeAvailabilityZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAvailabilityZoneResponse struct{}"
	}

	return strings.Join([]string{"ChangeAvailabilityZoneResponse", string(data)}, " ")
}
