package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTemplates 更新模板请求
type UpdateTemplates struct {

	// 模板名称,非必填
	Name *string `json:"name,omitempty"`

	// 模板描述,非必填
	Description *string `json:"description,omitempty"`

	// 模板标签,非必填
	Tags *[]string `json:"tags,omitempty"`

	// 模板状态,非必填
	Status *int32 `json:"status,omitempty"`

	// obs地址,必填
	ObsUrl *string `json:"obs_url,omitempty"`

	// 应用是否托管到servicestage,1是托管,0是不托管,非必填
	IsServicestage *int32 `json:"is_servicestage,omitempty"`
}

func (o UpdateTemplates) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTemplates struct{}"
	}

	return strings.Join([]string{"UpdateTemplates", string(data)}, " ")
}
