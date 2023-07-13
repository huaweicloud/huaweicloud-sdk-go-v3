package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubEnterpriseAmountRequest Request Object
type UpdateSubEnterpriseAmountRequest struct {
	Body *TransferEnterpriseMultiAccountReq `json:"body,omitempty"`
}

func (o UpdateSubEnterpriseAmountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubEnterpriseAmountRequest struct{}"
	}

	return strings.Join([]string{"UpdateSubEnterpriseAmountRequest", string(data)}, " ")
}
