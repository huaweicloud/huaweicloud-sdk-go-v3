package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KeystoneListMappingsResponse Response Object
type KeystoneListMappingsResponse struct {
	Links *Links `json:"links,omitempty"`

	// 映射信息列表。
	Mappings       *[]MappingResult `json:"mappings,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o KeystoneListMappingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListMappingsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListMappingsResponse", string(data)}, " ")
}
