package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoteMirrorSyncInfoDto struct {

	// **参数解释：** 用户名(需要base64加密)。
	Username *string `json:"username,omitempty"`

	// **参数解释：** 密码(需要base64加密)。
	Password *string `json:"password,omitempty"`

	// **参数解释：** 拓展点uuid。
	EndpointUuid *string `json:"endpoint_uuid,omitempty"`
}

func (o RemoteMirrorSyncInfoDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteMirrorSyncInfoDto struct{}"
	}

	return strings.Join([]string{"RemoteMirrorSyncInfoDto", string(data)}, " ")
}
