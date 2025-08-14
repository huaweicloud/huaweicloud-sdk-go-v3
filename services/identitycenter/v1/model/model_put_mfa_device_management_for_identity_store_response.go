package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PutMfaDeviceManagementForIdentityStoreResponse Response Object
type PutMfaDeviceManagementForIdentityStoreResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o PutMfaDeviceManagementForIdentityStoreResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutMfaDeviceManagementForIdentityStoreResponse struct{}"
	}

	return strings.Join([]string{"PutMfaDeviceManagementForIdentityStoreResponse", string(data)}, " ")
}
