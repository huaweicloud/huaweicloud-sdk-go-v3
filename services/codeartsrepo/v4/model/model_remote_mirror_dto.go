package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RemoteMirrorDto struct {

	// **参数解释：** 镜像地址。
	Url *string `json:"url,omitempty"`
}

func (o RemoteMirrorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteMirrorDto struct{}"
	}

	return strings.Join([]string{"RemoteMirrorDto", string(data)}, " ")
}
