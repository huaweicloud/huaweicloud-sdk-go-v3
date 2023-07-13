package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelAutoRenewalResourcesResponse Response Object
type CancelAutoRenewalResourcesResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelAutoRenewalResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelAutoRenewalResourcesResponse struct{}"
	}

	return strings.Join([]string{"CancelAutoRenewalResourcesResponse", string(data)}, " ")
}
