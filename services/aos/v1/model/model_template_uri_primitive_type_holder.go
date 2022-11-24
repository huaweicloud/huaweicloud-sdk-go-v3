package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateUriPrimitiveTypeHolder struct {

	// HCL模板的OBS地址，该模板描述了资源的目标状态  OBS地址必须为同region的OBS地址，暂不支持跨region访问  对应的文件应该是纯tf文件或zip压缩包  纯tf文件需要以`.tf`或者`.tfjson`结尾，并遵守hcl语法  压缩包目前只支持zip格式，文件需要以\".zip\"结尾。解压后的文件不得包含\".tfvars\"文件  template_body和template_uri 必须有且只有一个存在
	TemplateUri *string `json:"template_uri,omitempty"`
}

func (o TemplateUriPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateUriPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"TemplateUriPrimitiveTypeHolder", string(data)}, " ")
}
