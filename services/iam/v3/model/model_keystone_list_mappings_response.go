package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type KeystoneListMappingsResponse struct {
	Links *Links `json:"links,omitempty" xml:"links"`

	// 映射信息列表。
	Mappings       *[]MappingResult `json:"mappings,omitempty" xml:"mappings"`
	HttpStatusCode int              `json:"-"`
}

func (o KeystoneListMappingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListMappingsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListMappingsResponse", string(data)}, " ")
}
