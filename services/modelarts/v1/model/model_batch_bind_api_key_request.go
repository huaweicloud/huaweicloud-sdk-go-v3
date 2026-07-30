package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchBindApiKeyRequest struct {

	// **参数解释：** 请求批量绑定的api-key的id数组。 **约束限制：** 请求批量绑定api-key的id个数不超过10个。
	KeyIds []BatchBindApiKeyRequestKeyIds `json:"key_ids"`
}

func (o BatchBindApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchBindApiKeyRequest struct{}"
	}

	return strings.Join([]string{"BatchBindApiKeyRequest", string(data)}, " ")
}
