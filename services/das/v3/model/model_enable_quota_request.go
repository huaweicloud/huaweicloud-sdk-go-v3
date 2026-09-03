package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableQuotaRequest Request Object
type EnableQuotaRequest struct {
	Body *EnableQuotaRequestBody `json:"body,omitempty"`
}

func (o EnableQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableQuotaRequest struct{}"
	}

	return strings.Join([]string{"EnableQuotaRequest", string(data)}, " ")
}
