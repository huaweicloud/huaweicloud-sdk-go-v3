package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LandingZoneVersionDetail Landing Zone版本详细信息。
type LandingZoneVersionDetail struct {

	// Landing Zone版本。
	Version *string `json:"version,omitempty"`

	// Landing Zone版本发布时间。
	ReleaseDate *string `json:"release_date,omitempty"`

	// Landing Zone版本发布说明。
	ReleaseNotes *string `json:"release_notes,omitempty"`
}

func (o LandingZoneVersionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LandingZoneVersionDetail struct{}"
	}

	return strings.Join([]string{"LandingZoneVersionDetail", string(data)}, " ")
}
