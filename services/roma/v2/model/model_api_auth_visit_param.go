package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiAuthVisitParam struct {
	// 需要授权的API编号

	ApiId string `json:"api_id"`
	// 需要授权的APP编号

	AppId *string `json:"app_id,omitempty"`
	// 访问参数  支持英文、数字、下划线和中划线，多个参数以英文格式下的逗号隔开，单个参数须以英文或数字结尾且不能重复，且单个参数长度不超过255个字符。

	VisitParam string `json:"visit_param"`
}

func (o ApiAuthVisitParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiAuthVisitParam struct{}"
	}

	return strings.Join([]string{"ApiAuthVisitParam", string(data)}, " ")
}
