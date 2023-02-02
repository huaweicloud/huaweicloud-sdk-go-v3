package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VarsUriPrimitiveTypeHolder struct {

	// HCL支持参数，即，同一个模板可以给予不同的参数而达到不同的效果  * vars_uri需要指向一个OBS的pre-signed URL地址，其他地址暂不支持  * 资源编排服务支持vars_structure，vars_body和vars_uri，如果他们中声名了同一个变量，将报错400  * vars_uri中的内容使用HCL的tfvars格式，用户可以将”.tfvars”中的内容保存到文件并上传到OBS中，并将OBS pre-signed URL传递给vars_uri。具体tfvars格式见：https://www.terraform.io/language/values/variables#variable-definitions-tfvars-files  * 注意：vars_uri的内容不应该含有任何敏感信息，资源编排服务会直接明文使用、log、展示、存储对应的vars。如为敏感信息，建议通过vars_structure并设置encryption字段传递
	VarsUri *string `json:"vars_uri,omitempty"`
}

func (o VarsUriPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VarsUriPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"VarsUriPrimitiveTypeHolder", string(data)}, " ")
}
