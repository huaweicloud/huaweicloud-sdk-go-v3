package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRewriteUrlConfig 参数解释：URL重定向配置。  约束限制：当action为REDIRECT_TO_POOL的时候生效。
type CreateRewriteUrlConfig struct {

	// 参数解释：重定向的域名。  取值范围：英文字母、数字、“-”、“.”，必须以字母、数字开头。  默认取值：${host}，表示继承原值（即与被重写请求host保持一致）。
	Host *string `json:"host,omitempty"`

	// 参数解释：重定向的请求路径。其中$1-$9会匹配请求url通配符星号(*)，当正则匹配分组小于指定数字，则$指定数字结果为空。$后面跟字母，匹配结果均为空，直到下一个特殊字符出现，例如$abc#123，则匹配结果为#123；$后面跟特殊字符则直接输出特殊字符，例如$#匹配结果为$#。  取值范围：英文字母、数字、_~';@^-%#&$.+?,=!:|/()，且必须以\"/\"开头。  默认值${path}表示继承原值（即与被重写请求保持一致）。
	Path *string `json:"path,omitempty"`

	// 参数解释：重定向的查询字符串。其中$1-$9会匹配请求url通配符星号（*），当正则匹配分组小于指定数字，则$指定数字结果为空。$后面跟字母，匹配结果均为空，直到下一个特殊字符出现，例如$abc#123，则匹配结果为#123；$后面跟特殊字符则直接输出特殊字符，例如$#匹配结果为$#。  取值范围：英文字母、数字和特殊字符：!$&'()+,-./:;=?@^_`。字母区分大小写。  默认取值：${query}，表示继承原值（即与被重写请求保持一致）。
	Query *string `json:"query,omitempty"`
}

func (o CreateRewriteUrlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRewriteUrlConfig struct{}"
	}

	return strings.Join([]string{"CreateRewriteUrlConfig", string(data)}, " ")
}
