package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteConnectionState 实例状态。
type SiteConnectionState struct {
	State *SiteConnectionStateEnum `json:"state"`
}

func (o SiteConnectionState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteConnectionState struct{}"
	}

	return strings.Join([]string{"SiteConnectionState", string(data)}, " ")
}
