package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type EnableEnterpriseProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"EnableEnterpriseProjectResponse", string(data)}, " ")
}
