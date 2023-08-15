package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizePostPaidServerRequestBody This is a auto create Body Object
type ResizePostPaidServerRequestBody struct {
	Resize *ResizePostPaidServerOption `json:"resize"`

	// 是否只预检此次请求。  true：发送检查请求，不会变更云服务器规格。检查项包括是否填写了必需参数、请求格式等。  如果检查不通过，则返回对应错误。 如果检查通过，则返回202状态码。 false：发送正常请求，通过检查后并且执行变更云服务器规格请求。
	DryRun *bool `json:"dry_run,omitempty"`
}

func (o ResizePostPaidServerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizePostPaidServerRequestBody struct{}"
	}

	return strings.Join([]string{"ResizePostPaidServerRequestBody", string(data)}, " ")
}
