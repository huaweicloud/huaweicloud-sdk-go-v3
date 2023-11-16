package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateIpToPolicyResponse Response Object
type AssociateIpToPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateIpToPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateIpToPolicyResponse struct{}"
	}

	return strings.Join([]string{"AssociateIpToPolicyResponse", string(data)}, " ")
}
