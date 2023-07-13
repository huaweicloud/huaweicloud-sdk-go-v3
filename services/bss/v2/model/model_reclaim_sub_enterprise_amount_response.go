package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimSubEnterpriseAmountResponse Response Object
type ReclaimSubEnterpriseAmountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ReclaimSubEnterpriseAmountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimSubEnterpriseAmountResponse struct{}"
	}

	return strings.Join([]string{"ReclaimSubEnterpriseAmountResponse", string(data)}, " ")
}
