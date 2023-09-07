package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgenciesResponse Response Object
type ListAgenciesResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *AgencyKindObj `json:"kind,omitempty"`

	// 委托列表。
	Agencies       *[]AgencyItem `json:"agencies,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListAgenciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgenciesResponse struct{}"
	}

	return strings.Join([]string{"ListAgenciesResponse", string(data)}, " ")
}
