package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubEnterpriseAccountResponse Response Object
type CreateSubEnterpriseAccountResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSubEnterpriseAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubEnterpriseAccountResponse struct{}"
	}

	return strings.Join([]string{"CreateSubEnterpriseAccountResponse", string(data)}, " ")
}
