package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEnterpriseProjectResponse struct {
	EnterpriseProject *EpDetail `json:"enterprise_project,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o CreateEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"CreateEnterpriseProjectResponse", string(data)}, " ")
}
