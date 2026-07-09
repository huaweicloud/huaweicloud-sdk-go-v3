package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSwitchConfigInfoResponse Response Object
type ListSwitchConfigInfoResponse struct {
	SwitchInfo *SwitchInfo `json:"switch_info,omitempty"`

	VersionInfo    *VersionInfo `json:"version_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListSwitchConfigInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSwitchConfigInfoResponse struct{}"
	}

	return strings.Join([]string{"ListSwitchConfigInfoResponse", string(data)}, " ")
}
