package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetClusterQuotaRequest Request Object
type GetClusterQuotaRequest struct {
}

func (o GetClusterQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetClusterQuotaRequest struct{}"
	}

	return strings.Join([]string{"GetClusterQuotaRequest", string(data)}, " ")
}
