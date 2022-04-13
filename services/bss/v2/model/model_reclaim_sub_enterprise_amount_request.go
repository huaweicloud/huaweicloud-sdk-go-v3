package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
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
