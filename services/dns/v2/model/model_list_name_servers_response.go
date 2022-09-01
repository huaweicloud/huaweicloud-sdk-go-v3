package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNameServersResponse struct {
	Nameservers    *[]NameServersResp `json:"nameservers,omitempty" xml:"nameservers"`
	HttpStatusCode int                `json:"-"`
}

func (o ListNameServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNameServersResponse struct{}"
	}

	return strings.Join([]string{"ListNameServersResponse", string(data)}, " ")
}
