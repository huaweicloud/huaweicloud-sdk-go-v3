package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetHttpDetectResponseBodyDetail 终端探测详情
type GetHttpDetectResponseBodyDetail struct {

	// 探测终端返回的http返回码，0代表用户在黑名单无法发送，-1代表用户终端超过5秒未响应，-2代表队列已满，Http探测任务未执行。其他httpcode为终端实际返回值。
	HttpCode int32 `json:"httpCode"`

	// 终端探测响应体，如果httpCode为0，-1，-2, 2xx时响应体内容固定，由SMN定义。其余httpCode的响应体内容为终端返回内容。
	HttpResponse string `json:"httpResponse"`
}

func (o GetHttpDetectResponseBodyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHttpDetectResponseBodyDetail struct{}"
	}

	return strings.Join([]string{"GetHttpDetectResponseBodyDetail", string(data)}, " ")
}
