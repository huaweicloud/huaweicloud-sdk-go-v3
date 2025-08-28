package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVocabularySwitchConfigsResponse Response Object
type ShowVocabularySwitchConfigsResponse struct {

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 租户级的开关列表。
	Data           *[]SaveTtscTenantConfigsRequestBody `json:"data,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowVocabularySwitchConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVocabularySwitchConfigsResponse struct{}"
	}

	return strings.Join([]string{"ShowVocabularySwitchConfigsResponse", string(data)}, " ")
}
