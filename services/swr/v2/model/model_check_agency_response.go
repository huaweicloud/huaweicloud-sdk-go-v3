package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckAgencyResponse Response Object
type CheckAgencyResponse struct {

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 是否已创建委托
	IsAgency       *bool `json:"is_agency,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAgencyResponse struct{}"
	}

	return strings.Join([]string{"CheckAgencyResponse", string(data)}, " ")
}
