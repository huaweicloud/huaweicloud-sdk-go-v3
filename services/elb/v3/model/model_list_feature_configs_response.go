package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeatureConfigsResponse Response Object
type ListFeatureConfigsResponse struct {

	// **参数解释**：特性配置列表。
	Configs *[]FeatureConfig `json:"configs,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListFeatureConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeatureConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListFeatureConfigsResponse", string(data)}, " ")
}
