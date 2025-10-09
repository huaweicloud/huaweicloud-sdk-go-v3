package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetClusterQuotaResponse Response Object
type GetClusterQuotaResponse struct {
	Quotas         *GetClusterQuotaBodyQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o GetClusterQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetClusterQuotaResponse struct{}"
	}

	return strings.Join([]string{"GetClusterQuotaResponse", string(data)}, " ")
}
