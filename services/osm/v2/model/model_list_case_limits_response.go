package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCaseLimitsResponse struct {
	Config         *TenantConfigV2 `json:"config,omitempty" xml:"config"`
	HttpStatusCode int             `json:"-"`
}

func (o ListCaseLimitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCaseLimitsResponse struct{}"
	}

	return strings.Join([]string{"ListCaseLimitsResponse", string(data)}, " ")
}
