package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateDesktopsEipResponse Response Object
type AssociateDesktopsEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AssociateDesktopsEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateDesktopsEipResponse struct{}"
	}

	return strings.Join([]string{"AssociateDesktopsEipResponse", string(data)}, " ")
}
