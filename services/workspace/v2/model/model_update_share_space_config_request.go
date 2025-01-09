package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateShareSpaceConfigRequest Request Object
type UpdateShareSpaceConfigRequest struct {
	Body *UpdateShareSpaceConfigReq `json:"body,omitempty"`
}

func (o UpdateShareSpaceConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateShareSpaceConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateShareSpaceConfigRequest", string(data)}, " ")
}
