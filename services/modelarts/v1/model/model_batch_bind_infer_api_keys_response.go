package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchBindInferApiKeysResponse Response Object
type BatchBindInferApiKeysResponse struct {

	// **参数解释：** 请求绑定apikey总个数。 **取值范围：** 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** 绑定apikey成功个数。 **取值范围：** 不涉及。
	SuccessCount *int32 `json:"success_count,omitempty"`

	// **参数解释：** 绑定成功的apikey列表。
	SuccessItems *[]ApiKeyResponseV2 `json:"success_items,omitempty"`

	// **参数解释：** 绑定apikey失败个数。 **取值范围：** 不涉及。
	FailureCount *int32 `json:"failure_count,omitempty"`

	// **参数解释：** 绑定失败的apikey列表。
	FailureItems   *[]ApiKeyFailureResponse `json:"failure_items,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o BatchBindInferApiKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindInferApiKeysResponse struct{}"
	}

	return strings.Join([]string{"BatchBindInferApiKeysResponse", string(data)}, " ")
}
