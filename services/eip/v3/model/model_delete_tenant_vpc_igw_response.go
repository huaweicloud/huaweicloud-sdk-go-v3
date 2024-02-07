package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTenantVpcIgwResponse Response Object
type DeleteTenantVpcIgwResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteTenantVpcIgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTenantVpcIgwResponse struct{}"
	}

	return strings.Join([]string{"DeleteTenantVpcIgwResponse", string(data)}, " ")
}
