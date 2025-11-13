package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateOcspSwitchResponse Response Object
type UpdateOcspSwitchResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateOcspSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateOcspSwitchResponse struct{}"
	}

	return strings.Join([]string{"UpdateOcspSwitchResponse", string(data)}, " ")
}
