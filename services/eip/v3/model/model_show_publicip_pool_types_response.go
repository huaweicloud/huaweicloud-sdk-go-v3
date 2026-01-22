package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicipPoolTypesResponse Response Object
type ShowPublicipPoolTypesResponse struct {
	PublicipPoolTypes *PublicPoolType `json:"publicip-pool-types,omitempty"`
	HttpStatusCode    int             `json:"-"`
}

func (o ShowPublicipPoolTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicipPoolTypesResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicipPoolTypesResponse", string(data)}, " ")
}
