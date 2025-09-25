package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSlowlogSensitiveStatusResponse Response Object
type UpdateSlowlogSensitiveStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSlowlogSensitiveStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlowlogSensitiveStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateSlowlogSensitiveStatusResponse", string(data)}, " ")
}
