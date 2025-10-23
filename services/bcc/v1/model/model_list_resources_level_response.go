package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcesLevelResponse Response Object
type ListResourcesLevelResponse struct {

	// 资源分级列表数据
	ResourceLevelList *[]ResourceLevelItem `json:"resource_level_list,omitempty"`

	// 资源等级总数
	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListResourcesLevelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcesLevelResponse struct{}"
	}

	return strings.Join([]string{"ListResourcesLevelResponse", string(data)}, " ")
}
