package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReclaimSubEnterpriseAmountRequest Request Object
type ReclaimSubEnterpriseAmountRequest struct {
	Body *RetrieveEnterpriseMultiAccountReq `json:"body,omitempty"`
}

func (o ReclaimSubEnterpriseAmountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReclaimSubEnterpriseAmountRequest struct{}"
	}

	return strings.Join([]string{"ReclaimSubEnterpriseAmountRequest", string(data)}, " ")
}
