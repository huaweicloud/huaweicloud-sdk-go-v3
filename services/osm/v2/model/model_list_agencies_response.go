package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAgenciesResponse struct {
	// 总数

	TotalCount *int32 `json:"total_count,omitempty"`
	// 委托列表

	AgencyList     *[]AgencyV2 `json:"agency_list,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListAgenciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgenciesResponse struct{}"
	}

	return strings.Join([]string{"ListAgenciesResponse", string(data)}, " ")
}
