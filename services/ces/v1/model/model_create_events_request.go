package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEventsRequest struct {
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`
	// 上报自定义事件。请求参数。

	Body *[]EventItem `json:"body,omitempty"`
}

func (o CreateEventsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEventsRequest struct{}"
	}

	return strings.Join([]string{"CreateEventsRequest", string(data)}, " ")
}
