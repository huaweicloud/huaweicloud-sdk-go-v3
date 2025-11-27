package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncResourceAgentReq struct {

	// **参数解释：** 资源信息。 **约束限制：** 不涉及。 **取值范围：** 选择需要同步的资源对应的资源信息，列表大小1~100之间。 **默认取值：** 不涉及。
	ResourceInfos *[]SyncResourceAgentReqResourceInfos `json:"resource_infos,omitempty"`

	// **参数解释：** 云厂商。 **约束限制：** 不涉及。 **取值范围：** - 华为云资源同步时，可以不传。 - 阿里云资源同步时，传ALI。 **默认取值：** 不涉及。
	Vendor *string `json:"vendor,omitempty"`
}

func (o SyncResourceAgentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncResourceAgentReq struct{}"
	}

	return strings.Join([]string{"SyncResourceAgentReq", string(data)}, " ")
}
