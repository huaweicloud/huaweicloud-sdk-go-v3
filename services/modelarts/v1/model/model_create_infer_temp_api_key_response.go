package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInferTempApiKeyResponse Response Object
type CreateInferTempApiKeyResponse struct {

	// **参数解释：** 临时apikey。 **取值范围：**不涉及。
	ApiKey *string `json:"api_key,omitempty"`

	// **参数解释：** 临时apikey超时时间。 **取值范围：**不涉及。
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// **参数解释：** 临时apikey创建时间。 **取值范围：**不涉及。
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateInferTempApiKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInferTempApiKeyResponse struct{}"
	}

	return strings.Join([]string{"CreateInferTempApiKeyResponse", string(data)}, " ")
}
