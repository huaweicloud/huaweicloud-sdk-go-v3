package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpStatus **参数解释：** HTTP响应状态。 **取值范围：** - OK：请求已成功处理，服务器返回了请求的资源或操作结果。 - ACCEPTED：请求已被服务器接受，但处理尚未完成（通常用于异步操作）。 - INTERNAL_SERVER_ERROR：接口内部异常。 - BAD_REQUEST：参数校验异常。
type HttpStatus struct {
}

func (o HttpStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpStatus struct{}"
	}

	return strings.Join([]string{"HttpStatus", string(data)}, " ")
}
