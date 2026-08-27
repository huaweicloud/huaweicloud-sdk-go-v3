package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUnBindApiKeyRequestKeyIds struct {

	// **参数解释：** apikey_id，在创建API_KEY时即可在返回体中获取，也可通过查询api-keys列表获取当前用户拥有的apikey，其中key_id字段即为apikey_id。 **约束限制：** 不涉及。 **取值范围：** apikey_id只能由英文小写字母、数字组成，且长度为32个字符。 **默认取值：** 不涉及。
	KeyId *string `json:"key_id,omitempty"`
}

func (o BatchUnBindApiKeyRequestKeyIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUnBindApiKeyRequestKeyIds struct{}"
	}

	return strings.Join([]string{"BatchUnBindApiKeyRequestKeyIds", string(data)}, " ")
}
