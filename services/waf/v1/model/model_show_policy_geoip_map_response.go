package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPolicyGeoipMapResponse Response Object
type ShowPolicyGeoipMapResponse struct {

	// **参数解释：** 各个洲上的国家名称分布信息 **取值范围：** 不涉及
	Continent *interface{} `json:"continent,omitempty"`

	// **参数解释：** key值为各个国家的简称（AB和AB2除外，AB表示海外及港澳台，AB2表示海外），当key为CN时，里面是数组内容为各个省份的简称 **取值范围：** 不涉及
	Geomap *interface{} `json:"geomap,omitempty"`

	// **参数解释：** geomap中的值对应语言的显示名称 **取值范围：** 不涉及
	Locale         *interface{} `json:"locale,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowPolicyGeoipMapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPolicyGeoipMapResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyGeoipMapResponse", string(data)}, " ")
}
