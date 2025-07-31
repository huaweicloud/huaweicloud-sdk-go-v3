package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicZoneNameServerResponse Response Object
type ShowPublicZoneNameServerResponse struct {

	// **参数解释：** 查询公网域名的名称服务器响应。 **取值范围：** 不涉及。
	Nameservers    *[]Nameserver `json:"nameservers,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowPublicZoneNameServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicZoneNameServerResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicZoneNameServerResponse", string(data)}, " ")
}
