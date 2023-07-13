package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStorageClaimResponse Response Object
type DeleteStorageClaimResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteStorageClaimResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStorageClaimResponse struct{}"
	}

	return strings.Join([]string{"DeleteStorageClaimResponse", string(data)}, " ")
}
