package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateZoneNameServerResponse Response Object
type ShowPrivateZoneNameServerResponse struct {

	// **参数解释：** 名称服务器列表信息。 **取值范围：** 不涉及。
	Nameservers    *[]PrivateNameServer `json:"nameservers,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowPrivateZoneNameServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateZoneNameServerResponse struct{}"
	}

	return strings.Join([]string{"ShowPrivateZoneNameServerResponse", string(data)}, " ")
}
