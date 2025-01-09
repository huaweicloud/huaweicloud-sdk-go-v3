package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEnterpriseIdRequest Request Object
type UpdateEnterpriseIdRequest struct {
	Body *ModifyEnterpriseIdReq `json:"body,omitempty"`
}

func (o UpdateEnterpriseIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEnterpriseIdRequest struct{}"
	}

	return strings.Join([]string{"UpdateEnterpriseIdRequest", string(data)}, " ")
}
