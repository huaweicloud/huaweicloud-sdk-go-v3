package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUnBindApiKeyRequest struct {

	// **参数解释：** 请求批量解绑的api-key的id数组。 **约束限制：** 请求批量解绑api-key的id个数不超过10个。
	KeyIds []BatchUnBindApiKeyRequestKeyIds `json:"key_ids"`
}

func (o BatchUnBindApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnBindApiKeyRequest struct{}"
	}

	return strings.Join([]string{"BatchUnBindApiKeyRequest", string(data)}, " ")
}
