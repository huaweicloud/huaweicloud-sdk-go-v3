package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModuleUriPrimitiveTypeHolder struct {

	// 模块（module）包的OBS地址。模块允许用户将可复用的代码编辑在一起供模块使用。  OBS地址支持同类型Region之间进行互相访问（Region分为通用Region和专属Region，通用Region指面向公共租户提供通用云服务的Region；专属Region指只承载同一类业务或只面向特定租户提供业务服务的专用Region）  模块包只支持zip压缩包，文件需要以\".zip\"结尾。关于模块包的校验要求如下：   * 模块包中不得包含以\".tfvars\"结尾的文件   * 模块包解压前后的大小均应控制在1MB以内   * 模块包内的文件数量不能超过100个   * 模块包内的文件路径不允许以正斜线（/）开头   * 模块包内的文件路径分隔符之间不允许为空、`.`和`..`   * 模块包内的文件路径最长为2048   * 模块包内的文件名最长为255   * 模块包内应至少有一份模板文件（以`.tf`或`.tf.json`结尾的文件）  **注意：**   * 模块中的内容不支持敏感数据加密，资源编排服务会直接明文使用、log、展示、存储module_uri对应的模块包。
	ModuleUri *string `json:"module_uri,omitempty"`
}

func (o ModuleUriPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleUriPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ModuleUriPrimitiveTypeHolder", string(data)}, " ")
}
