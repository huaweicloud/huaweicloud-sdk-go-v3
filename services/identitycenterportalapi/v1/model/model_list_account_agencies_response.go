package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccountAgenciesResponse Response Object
type ListAccountAgenciesResponse struct {

	// 满足查询条件的委托或信任委托对象列表
	AgencyList *[]AgencyInfo `json:"agency_list,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListAccountAgenciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccountAgenciesResponse struct{}"
	}

	return strings.Join([]string{"ListAccountAgenciesResponse", string(data)}, " ")
}
