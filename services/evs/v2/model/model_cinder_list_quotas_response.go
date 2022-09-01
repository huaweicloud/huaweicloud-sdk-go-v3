package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CinderListQuotasResponse struct {
	QuotaSet       *QuotaList `json:"quota_set,omitempty" xml:"quota_set"`
	HttpStatusCode int        `json:"-"`
}

func (o CinderListQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListQuotasResponse struct{}"
	}

	return strings.Join([]string{"CinderListQuotasResponse", string(data)}, " ")
}
