package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTemplate 创建模板请求体
type CreateTemplate struct {

	// 模板名称,必填
	Name string `json:"name"`

	// 应用是否托管到servicestage,1是托管,0是不托管,非必填
	IsServicestage *int32 `json:"is_servicestage,omitempty"`

	// 描述信息,非必填
	Desc *string `json:"desc,omitempty"`

	// 模板在桶中的url,必填
	ObsUrl string `json:"obs_url"`

	// 模板标签,非必填
	Tags *[]string `json:"tags,omitempty"`
}

func (o CreateTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTemplate struct{}"
	}

	return strings.Join([]string{"CreateTemplate", string(data)}, " ")
}
