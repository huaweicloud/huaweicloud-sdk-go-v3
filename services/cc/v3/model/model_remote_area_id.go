package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoteAreaId 对端大区ID。
type RemoteAreaId struct {
	RemoteAreaId *RemoteAreaIdDef `json:"remote_area_id"`
}

func (o RemoteAreaId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoteAreaId struct{}"
	}

	return strings.Join([]string{"RemoteAreaId", string(data)}, " ")
}
