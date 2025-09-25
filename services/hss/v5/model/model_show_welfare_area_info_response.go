package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWelfareAreaInfoResponse Response Object
type ShowWelfareAreaInfoResponse struct {
	HotInfo *HttpWelfareInfoResponseInfoHotInfo `json:"hot_info,omitempty"`

	VersionUpdateInfo *HttpWelfareInfoResponseInfoVersionUpdateInfo `json:"version_update_info,omitempty"`

	ActivitiesInfo *HttpWelfareInfoResponseInfoActivitiesInfo `json:"activities_info,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o ShowWelfareAreaInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWelfareAreaInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowWelfareAreaInfoResponse", string(data)}, " ")
}
