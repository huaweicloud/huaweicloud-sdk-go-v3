package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FunctionGraphContentRsp SMN数据源配置内容
type FunctionGraphContentRsp struct {

	// 名称
	FunctionName *string `json:"functionName,omitempty"`

	// 原始URL
	OrigUrl *string `json:"origUrl,omitempty"`

	// 转换后的URL
	FinalUrl *string `json:"finalUrl,omitempty"`

	// 校验参数
	VerifyBody *string `json:"verifyBody,omitempty"`

	// 租户的AK
	Ak *string `json:"ak,omitempty"`

	// 租户的SK
	Sk *string `json:"sk,omitempty"`

	// 项目id
	ProjectId *string `json:"projectId,omitempty"`
}

func (o FunctionGraphContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionGraphContentRsp struct{}"
	}

	return strings.Join([]string{"FunctionGraphContentRsp", string(data)}, " ")
}
