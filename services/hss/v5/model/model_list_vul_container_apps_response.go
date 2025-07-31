package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulContainerAppsResponse Response Object
type ListVulContainerAppsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 软件信息列表
	DataList       *[]VulAffectContainerAppInfo `json:"data_list,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListVulContainerAppsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulContainerAppsResponse struct{}"
	}

	return strings.Join([]string{"ListVulContainerAppsResponse", string(data)}, " ")
}
