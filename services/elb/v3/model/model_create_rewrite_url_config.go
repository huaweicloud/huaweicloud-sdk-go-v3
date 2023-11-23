package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRewriteUrlConfig 转发到的后端主机组的URL配置。rewrite_url_enable为true时，该字段必须传入。
type CreateRewriteUrlConfig struct {

	// url重写的主机名。字符串只能包含英文字母、数字、“-”、“.”，必须以字母、数字开头。默认值${host}表示继承原值（即与被重写请求保持一致）。
	Host *string `json:"host,omitempty"`

	// url重定向的路径。默认值${path}表示继承原值（即与被重写请求保持一致）。 只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:|\\/()，且必须以\"/\"开头。其中$1-$9会匹配请求url通配符星号(*)，当正则匹配分组小于指定数字，则$指定数字结果为空。$后面跟字母，匹配结果均为空，直到下一个特殊字符出现，例如$abc#123，则匹配结果为#123；$后面跟特殊字符则直接输出特殊字符，例如$#匹配结果为$#.
	Path *string `json:"path,omitempty"`

	// url重定向的查询字符串。默认${query}表示继承原值（即与被重写请求保持一致）。 只能包含英文字母、数字和特殊字符：!$&'()*+,-./:;=?@^_`。字母区分大小写。其中$1-$9会匹配请求url通配符星号（*），当正则匹配分组小于指定数字，则$指定数字结果为空。$后面跟字母，匹配结果均为空，直到下一个特殊字符出现，例如$abc#123，则匹配结果为#123；$后面跟特殊字符则直接输出特殊字符，例如$#匹配结果为$#
	Query *string `json:"query,omitempty"`
}

func (o CreateRewriteUrlConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRewriteUrlConfig struct{}"
	}

	return strings.Join([]string{"CreateRewriteUrlConfig", string(data)}, " ")
}
