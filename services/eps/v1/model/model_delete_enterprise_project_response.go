package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteEnterpriseProjectResponse Response Object
type DeleteEnterpriseProjectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"DeleteEnterpriseProjectResponse", string(data)}, " ")
}
