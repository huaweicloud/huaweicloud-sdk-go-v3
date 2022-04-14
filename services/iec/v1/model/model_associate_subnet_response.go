package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AssociateSubnetResponse struct {
	Routetable     *Routetable `json:"routetable,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o AssociateSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSubnetResponse struct{}"
	}

	return strings.Join([]string{"AssociateSubnetResponse", string(data)}, " ")
}
