package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulAffectImageAppsResponse Response Object
type ListVulAffectImageAppsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 列表
	DataList       *[]VulAffectImageAppsResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ListVulAffectImageAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulAffectImageAppsResponse struct{}"
	}

	return strings.Join([]string{"ListVulAffectImageAppsResponse", string(data)}, " ")
}
