package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateResourceShareResponse Response Object
type DisassociateResourceShareResponse struct {
	ResourceShareAssociations *[]ResourceShareAssociation `json:"resource_share_associations,omitempty"`
	HttpStatusCode            int                         `json:"-"`
}

func (o DisassociateResourceShareResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateResourceShareResponse struct{}"
	}

	return strings.Join([]string{"DisassociateResourceShareResponse", string(data)}, " ")
}
