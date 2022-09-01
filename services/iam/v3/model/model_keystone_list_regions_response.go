package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type KeystoneListRegionsResponse struct {
	Links *Links `json:"links,omitempty" xml:"links"`

	// 区域信息列表。
	Regions        *[]Region `json:"regions,omitempty" xml:"regions"`
	HttpStatusCode int       `json:"-"`
}

func (o KeystoneListRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneListRegionsResponse struct{}"
	}

	return strings.Join([]string{"KeystoneListRegionsResponse", string(data)}, " ")
}
