package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SyncReq struct {
	Sync *SyncParam `json:"sync"`
}

func (o SyncReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SyncReq struct{}"
	}

	return strings.Join([]string{"SyncReq", string(data)}, " ")
}
