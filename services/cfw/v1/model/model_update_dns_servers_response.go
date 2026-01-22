package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDnsServersResponse Response Object
type UpdateDnsServersResponse struct {

	// **参数解释**： 域名服务器列表 **取值范围**： 不涉及
	Data           *[]string `json:"data,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateDnsServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDnsServersResponse struct{}"
	}

	return strings.Join([]string{"UpdateDnsServersResponse", string(data)}, " ")
}
