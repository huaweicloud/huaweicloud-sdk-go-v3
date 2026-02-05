package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiApiKeysResponse Response Object
type ListAiApiKeysResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// AIAPIKey列表，展示匿名化的已经创建的AIAPIKey。 AIAPIKey的创建数量上限可以通过配额调整。
	AiApiKeys      *[]AiApiKeyBaseInfo `json:"ai_api_keys,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListAiApiKeysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiApiKeysResponse struct{}"
	}

	return strings.Join([]string{"ListAiApiKeysResponse", string(data)}, " ")
}
