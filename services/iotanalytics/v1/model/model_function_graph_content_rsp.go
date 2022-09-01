package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SMN数据源配置内容
type FunctionGraphContentRsp struct {

	// 名称
	FunctionName *string `json:"functionName,omitempty" xml:"functionName"`

	// 原始URL
	OrigUrl *string `json:"origUrl,omitempty" xml:"origUrl"`

	// 转换后的URL
	FinalUrl *string `json:"finalUrl,omitempty" xml:"finalUrl"`

	// 校验参数
	VerifyBody *string `json:"verifyBody,omitempty" xml:"verifyBody"`

	// 租户的AK
	Ak *string `json:"ak,omitempty" xml:"ak"`

	// 租户的SK
	Sk *string `json:"sk,omitempty" xml:"sk"`

	// 项目id
	ProjectId *string `json:"projectId,omitempty" xml:"projectId"`
}

func (o FunctionGraphContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionGraphContentRsp struct{}"
	}

	return strings.Join([]string{"FunctionGraphContentRsp", string(data)}, " ")
}
