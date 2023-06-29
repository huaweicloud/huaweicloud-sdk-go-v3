package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubEnterpriseAmountResponse Response Object
type UpdateSubEnterpriseAmountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSubEnterpriseAmountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubEnterpriseAmountResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubEnterpriseAmountResponse", string(data)}, " ")
}
