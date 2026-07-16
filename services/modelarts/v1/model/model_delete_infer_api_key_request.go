package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInferApiKeyRequest Request Object
type DeleteInferApiKeyRequest struct {

	// **参数解释：** apikey_id，在[创建API_KEY](CreateInferApiKey.xml)时即可在返回体中获取，也可通过[查询api-keys列表](ListInferApiKeys.xml)获取当前用户拥有的apikey，其中key_id字段即为apikey_id。 **约束限制：** 不涉及。 **取值范围：** apikey_id只能由英文小写字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及。
	KeyId string `json:"key_id"`
}

func (o DeleteInferApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInferApiKeyRequest struct{}"
	}

	return strings.Join([]string{"DeleteInferApiKeyRequest", string(data)}, " ")
}
