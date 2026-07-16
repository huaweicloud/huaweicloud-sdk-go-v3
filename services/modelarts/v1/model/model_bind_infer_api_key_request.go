package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindInferApiKeyRequest Request Object
type BindInferApiKeyRequest struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及 **取值范围：** 服务ID **默认取值：** 不涉及
	ServiceId string `json:"service_id"`

	// **参数解释：** apikey_id，在[创建API_KEY](CreateInferApiKey.xml)时即可在返回体中获取，也可通过[查询api-keys列表](ListInferApiKeys.xml)获取当前用户拥有的apikey，其中key_id字段即为apikey_id。 **约束限制：** 不涉及 **取值范围：** apikey_id只能由英文小写字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及
	KeyId string `json:"key_id"`
}

func (o BindInferApiKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindInferApiKeyRequest struct{}"
	}

	return strings.Join([]string{"BindInferApiKeyRequest", string(data)}, " ")
}
