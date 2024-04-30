package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateUriPrimitiveTypeHolder struct {

	// HCL模板的OBS地址，该模板描述了资源的目标状态。资源编排服务将比较此模板与当前远程资源的状态之间的区别  OBS地址支持同类型Region之间进行互相访问（Region分为通用Region和专属Region，通用Region指面向公共租户提供通用云服务的Region；专属Region指只承载同一类业务或只面向特定租户提供业务服务的专用Region）  对应的文件应该是纯tf文件或zip压缩包  纯tf文件需要以“.tf”或者“.tf.json”结尾，并遵守HCL语法，且文件的编码格式须为UTF-8  压缩包目前只支持zip格式，文件需要以\".zip\"结尾。解压后的文件不得包含\".tfvars\"文件。解压前最大支持1MB，解压后最大支持1MB。zip压缩包文件数量不能超过100个  template_body和template_uri 必须有且只有一个存在  **注意：**   * 资源栈集不支持敏感数据加密，资源编排服务会直接明文使用、log、展示、存储template_uri对应的模板文件内容。   * template_uri对应的模板文件如果为zip类型，则内部的文件或文件夹名称长度不得超过255个字节，最深路径长度不得超过2048字节，zip包大小不得超过1MB
	TemplateUri *string `json:"template_uri,omitempty"`
}

func (o TemplateUriPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateUriPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"TemplateUriPrimitiveTypeHolder", string(data)}, " ")
}
