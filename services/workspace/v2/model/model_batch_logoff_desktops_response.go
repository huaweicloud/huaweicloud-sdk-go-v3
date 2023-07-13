package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchLogoffDesktopsResponse Response Object
type BatchLogoffDesktopsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o BatchLogoffDesktopsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchLogoffDesktopsResponse struct{}"
	}

	return strings.Join([]string{"BatchLogoffDesktopsResponse", string(data)}, " ")
}
