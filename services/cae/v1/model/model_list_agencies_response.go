package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgenciesResponse Response Object
type ListAgenciesResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“Agency”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

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
