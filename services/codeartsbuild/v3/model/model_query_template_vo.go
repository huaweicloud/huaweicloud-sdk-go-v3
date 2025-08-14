package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryTemplateVo 模板中构建执行步骤列表
type QueryTemplateVo struct {

	// 构建执行的步骤。
	Steps *[]CreateBuildJobSteps `json:"steps,omitempty"`

	// 构建步骤中的action。
	Actions *interface{} `json:"actions,omitempty"`

	// 是否自动更新子模块。
	AutoUpdateSubModule *bool `json:"auto_update_sub_module,omitempty"`

	// 配置镜像地址。
	Image *string `json:"image,omitempty"`

	// 配置镜像源的地址。
	ImageSource *string `json:"image_source,omitempty"`
}

func (o QueryTemplateVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTemplateVo struct{}"
	}

	return strings.Join([]string{"QueryTemplateVo", string(data)}, " ")
}
