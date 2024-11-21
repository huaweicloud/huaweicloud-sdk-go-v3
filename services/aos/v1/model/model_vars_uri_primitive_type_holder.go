package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VarsUriPrimitiveTypeHolder struct {

	// HCL参数文件的OBS地址。HCL模板支持参数传入，即，同一个模板可以给予不同的参数而达到不同的效果。  请确保OBS地址所在局点与使用RFS服务局点一致。  * vars_uri需要指向一个OBS的pre-signed URL地址，其他地址暂不支持  * 资源编排服务支持vars_body和vars_uri，如果以上两种方式中声明了同一个变量，将报错400  * vars_uri中的内容使用HCL的tfvars格式，用户可以将“.tfvars”中的内容保存到文件并上传到OBS中，并将OBS pre-signed URL传递给vars_uri  * 资源栈集不支持敏感数据加密，资源编排服务会直接明文使用、log、展示、存储vars_uri对应的参数文件内容
	VarsUri *string `json:"vars_uri,omitempty"`
}

func (o VarsUriPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VarsUriPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"VarsUriPrimitiveTypeHolder", string(data)}, " ")
}
