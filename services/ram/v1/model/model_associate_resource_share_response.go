package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AssociateResourceShareResponse struct {
	ResourceShareAssociations *[]ResourceShareAssociation `json:"resource_share_associations,omitempty"`
	HttpStatusCode            int                         `json:"-"`
}

func (o AssociateResourceShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateResourceShareResponse struct{}"
	}

	return strings.Join([]string{"AssociateResourceShareResponse", string(data)}, " ")
}
