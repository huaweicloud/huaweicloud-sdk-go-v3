package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
