package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataDetailRequest Request Object
type ShowDataDetailRequest struct {

	// 实例id
	Instance string `json:"instance"`

	// 资产guid
	Guid string `json:"guid"`

	// 是否忽略关联资产 缺省值：false
	IgnoreRelationships *bool `json:"ignore_relationships,omitempty"`
}

func (o ShowDataDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowDataDetailRequest", string(data)}, " ")
}
