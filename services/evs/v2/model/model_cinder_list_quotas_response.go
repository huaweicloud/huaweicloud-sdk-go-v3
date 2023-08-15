package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CinderListQuotasResponse Response Object
type CinderListQuotasResponse struct {
	QuotaSet       *QuotaList `json:"quota_set,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o CinderListQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CinderListQuotasResponse struct{}"
	}

	return strings.Join([]string{"CinderListQuotasResponse", string(data)}, " ")
}
