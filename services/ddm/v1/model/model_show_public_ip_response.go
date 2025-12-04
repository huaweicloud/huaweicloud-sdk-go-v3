package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPublicIpResponse Response Object
type ShowPublicIpResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 弹性公网IP信息。
	PublicIpList   *[]PublicIpInfo `json:"public_ip_list,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowPublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicIpResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicIpResponse", string(data)}, " ")
}
