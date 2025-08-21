package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryWebhookLogResponse Response Object
type ShowRepositoryWebhookLogResponse struct {

	// **参数解释：** Webhook 日志id。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** Webhook id。
	WebHookId *int32 `json:"web_hook_id,omitempty"`

	// **参数解释：** 触发类型。
	Trigger *string `json:"trigger,omitempty"`

	// **参数解释：** 请求地址。
	Url *string `json:"url,omitempty"`

	// **参数解释：** 响应状态，默认是响应码，如果webhook地址未返回或者其他异常情况，则记录为internal error
	ResponseStatus *string `json:"response_status,omitempty"`

	// **参数解释：** 响应耗时，单位是秒
	ExecutionDuration *float64 `json:"execution_duration,omitempty"`

	// **参数解释：** Webhook每次发送消息的时候，会在请求体中带上uuid字段，此处为该记录的uuid字段
	Uuid *string `json:"uuid,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 请求头，此处会将敏感信息如传递的token隐藏
	RequestHeaders *interface{} `json:"request_headers,omitempty"`

	// **参数解释：** 请求体，此处会将用户邮箱隐藏
	RequestData *interface{} `json:"request_data,omitempty"`

	// **参数解释：** 响应头
	ResponseHeaders *interface{} `json:"response_headers,omitempty"`

	// **参数解释：** 响应体
	ResponseBody   *interface{} `json:"response_body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowRepositoryWebhookLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryWebhookLogResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryWebhookLogResponse", string(data)}, " ")
}
