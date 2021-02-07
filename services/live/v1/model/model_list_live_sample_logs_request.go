package model

import (
	"encoding/json"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"

	"strings"
)

// Request Object
type ListLiveSampleLogsRequest struct {
	PlayDomain string           `json:"play_domain"`
	StartTime  *sdktime.SdkTime `json:"start_time"`
	EndTime    *sdktime.SdkTime `json:"end_time"`
}

func (o ListLiveSampleLogsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListLiveSampleLogsRequest struct{}"
	}

	return strings.Join([]string{"ListLiveSampleLogsRequest", string(data)}, " ")
}
