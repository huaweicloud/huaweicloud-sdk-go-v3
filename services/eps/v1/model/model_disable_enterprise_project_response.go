package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableEnterpriseProjectResponse Response Object
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
