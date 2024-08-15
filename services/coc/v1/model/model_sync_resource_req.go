package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncResourceReq struct {

	// 资源提供者
	Provider string `json:"provider"`

	// 资源类型
	Type string `json:"type"`
}

func (o SyncResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncResourceReq struct{}"
	}

	return strings.Join([]string{"SyncResourceReq", string(data)}, " ")
}
