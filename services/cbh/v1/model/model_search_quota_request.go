package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchQuotaRequest Request Object
type SearchQuotaRequest struct {
}

func (o SearchQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQuotaRequest struct{}"
	}

	return strings.Join([]string{"SearchQuotaRequest", string(data)}, " ")
}
