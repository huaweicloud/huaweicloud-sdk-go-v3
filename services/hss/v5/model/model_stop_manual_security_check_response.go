package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopManualSecurityCheckResponse Response Object
type StopManualSecurityCheckResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopManualSecurityCheckResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopManualSecurityCheckResponse struct{}"
	}

	return strings.Join([]string{"StopManualSecurityCheckResponse", string(data)}, " ")
}
