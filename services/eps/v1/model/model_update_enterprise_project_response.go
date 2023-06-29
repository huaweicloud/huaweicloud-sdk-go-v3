package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnterpriseProjectResponse Response Object
type UpdateEnterpriseProjectResponse struct {
	EnterpriseProject *EpDetail `json:"enterprise_project,omitempty"`
	HttpStatusCode    int       `json:"-"`
}

func (o UpdateEnterpriseProjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseProjectResponse struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseProjectResponse", string(data)}, " ")
}
