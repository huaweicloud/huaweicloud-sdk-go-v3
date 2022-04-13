package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FunctionGraph数据源配置内容
type FunctionGraphContentReq struct {
	// 名称

	FunctionName string `json:"function_name"`
	// 原始URL

	OrigUrl string `json:"orig_url"`
	// 转换后的URL

	FinalUrl string `json:"final_url"`
	// 校验参数

	VerifyBody string `json:"verify_body"`
	// 租户的AK

	Ak string `json:"ak"`
	// 租户的SK

	Sk string `json:"sk"`
	// 项目id

	ProjectId string `json:"project_id"`
}

func (o FunctionGraphContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionGraphContentReq struct{}"
	}

	return strings.Join([]string{"FunctionGraphContentReq", string(data)}, " ")
}
