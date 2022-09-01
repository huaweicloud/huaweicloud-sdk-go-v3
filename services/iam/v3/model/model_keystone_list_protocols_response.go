package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type KeystoneListProtocolsResponse struct {
	Links *Links `json:"links,omitempty" xml:"links"`

	// 协议信息列表。
	Protocols      *[]ProtocolResult `json:"protocols,omitempty" xml:"protocols"`
	HttpStatusCode int               `json:"-"`
}

func (o KeystoneListProtocolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListProtocolsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListProtocolsResponse", string(data)}, " ")
}
