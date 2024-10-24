package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnvsRequest Request Object
type ListEnvsRequest struct {

	// 页码
	PageNum *int64 `json:"page_num,omitempty"`

	// 当前页大小
	PageSize *int64 `json:"page_size,omitempty"`

	// 云服务类型 - STUDIO：设计态服务。 - CLOUD_BASIC：公有云基础版数据建模引擎。 - CLOUD_TRIAL：公有云体验版数据建模引擎。 - EDGE_BASIC：边缘云基础版数据建模引擎。 - CLOUD_LINKX：公有云基础版数字主线引擎。 - EDGE_LINKX：边缘云基础版数字主线引擎。
	EnvTypes *string `json:"env_types,omitempty"`
}

func (o ListEnvsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnvsRequest struct{}"
	}

	return strings.Join([]string{"ListEnvsRequest", string(data)}, " ")
}
