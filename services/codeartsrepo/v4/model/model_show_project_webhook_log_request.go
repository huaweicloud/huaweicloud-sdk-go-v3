package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectWebhookLogRequest Request Object
type ShowProjectWebhookLogRequest struct {

	// **参数解释：** 项目的32位uuid，项目唯一标识，通过[[查询项目列表](https://support.huaweicloud.com/api-projectman/ListProjectsV4.html)](tag:hws)[[查询项目列表](https://support.huaweicloud.com/intl/en-us/api-projectman/ListProjectsV4.html)](tag:hws_hk)[查询项目列表](tag:hcs,hcs_sm)接口查询项目列表获取。 **取值范围：** 字符串长度32。
	ProjectId string `json:"project_id"`

	// **参数解释：**  Webhook id。
	HookId int32 `json:"hook_id"`

	// **参数解释：**  Webhook 日志id。
	LogId int32 `json:"log_id"`
}

func (o ShowProjectWebhookLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectWebhookLogRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectWebhookLogRequest", string(data)}, " ")
}
