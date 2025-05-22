package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadHotKeyRequest Request Object
type DownloadHotKeyRequest struct {

	// **参数解释**： 实例ID。可通过DCS控制台进入实例详情界面查看。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 热Key分析任务ID，可通过DCS控制台进入实例缓存分析界面查看。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TaskId string `json:"task_id"`
}

func (o DownloadHotKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadHotKeyRequest struct{}"
	}

	return strings.Join([]string{"DownloadHotKeyRequest", string(data)}, " ")
}
