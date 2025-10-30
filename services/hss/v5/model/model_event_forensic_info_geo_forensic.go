package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventForensicInfoGeoForensic **参数解释**： 地理位置取证信息 **取值范围**： 不涉及
type EventForensicInfoGeoForensic struct {

	// **参数解释**： 源国家 **取值范围**： 不涉及
	SrcCountry *string `json:"src_country,omitempty"`

	// **参数解释**： 源城市 **取值范围**： 不涉及
	SrcCity *string `json:"src_city,omitempty"`

	// **参数解释**： 源纬度 **取值范围**： 不涉及
	SrcLatitude *int32 `json:"src_latitude,omitempty"`

	// **参数解释**： 源经度 **取值范围**： 不涉及
	SrcLongitude *int32 `json:"src_longitude,omitempty"`

	// **参数解释**： 目的国家 **取值范围**： 不涉及
	DestCountry *string `json:"dest_country,omitempty"`

	// **参数解释**： 目的城市 **取值范围**： 不涉及
	DestCity *string `json:"dest_city,omitempty"`

	// **参数解释**： 目的纬度 **取值范围**： 不涉及
	DestLatitude *int32 `json:"dest_latitude,omitempty"`

	// **参数解释**： 目的经度 **取值范围**： 不涉及
	DestLongitude *int32 `json:"dest_longitude,omitempty"`
}

func (o EventForensicInfoGeoForensic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventForensicInfoGeoForensic struct{}"
	}

	return strings.Join([]string{"EventForensicInfoGeoForensic", string(data)}, " ")
}
