package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FunctionGraph数据源配置内容
type FunctionGraphContentReq struct {

	// 名称
	FunctionName string `json:"function_name" xml:"function_name"`

	// 原始URL
	OrigUrl string `json:"orig_url" xml:"orig_url"`

	// 转换后的URL
	FinalUrl string `json:"final_url" xml:"final_url"`

	// 校验参数
	VerifyBody string `json:"verify_body" xml:"verify_body"`

	// 租户的AK
	Ak string `json:"ak" xml:"ak"`

	// 租户的SK
	Sk string `json:"sk" xml:"sk"`

	// 项目id
	ProjectId string `json:"project_id" xml:"project_id"`
}

func (o FunctionGraphContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionGraphContentReq struct{}"
	}

	return strings.Join([]string{"FunctionGraphContentReq", string(data)}, " ")
}
