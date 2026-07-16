package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePodIdentityAssociationResponse Response Object
type DeletePodIdentityAssociationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePodIdentityAssociationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePodIdentityAssociationResponse struct{}"
	}

	return strings.Join([]string{"DeletePodIdentityAssociationResponse", string(data)}, " ")
}
