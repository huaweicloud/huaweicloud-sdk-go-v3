package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoteMirrorSyncInfoDto struct {

	// **参数解释：** 用户名(需要base64加密)。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及
	Username *string `json:"username,omitempty"`

	// **参数解释：** 密码(需要base64加密)。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及
	Password *string `json:"password,omitempty"`

	// **参数解释：** 拓展点uuid。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	EndpointUuid *string `json:"endpoint_uuid,omitempty"`

	// **参数解释：** 允许强制同步。 **约束限制：** 不涉及。 **取值范围：** - true，强制同步。 - false，不强制同步。 **默认取值：** 不涉及。
	ForceFetch *bool `json:"force_fetch,omitempty"`
}

func (o RemoteMirrorSyncInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteMirrorSyncInfoDto struct{}"
	}

	return strings.Join([]string{"RemoteMirrorSyncInfoDto", string(data)}, " ")
}
