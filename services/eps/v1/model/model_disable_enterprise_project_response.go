package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DisableEnterpriseProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"DisableEnterpriseProjectResponse", string(data)}, " ")
}
