package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloneListenerResponse Response Object
type CloneListenerResponse struct {

	// **参数解释**：新监听器相关信息。
	ListenerList *[]CloneListenerResp `json:"listener_list,omitempty"`

	// **参数解释**：请求的ID 。 **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId *string `json:"request_id,omitempty"`

	// **参数解释**：监听器复制任务的ID，任务详情可通过GET /v3/{project_id}/elb/jobs/{job_id}进行查询。 **取值范围**：标准的UUID格式，长度为36个字符。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CloneListenerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneListenerResponse struct{}"
	}

	return strings.Join([]string{"CloneListenerResponse", string(data)}, " ")
}
