package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiKeyFailureResponse apikey响应信息
type ApiKeyFailureResponse struct {

	// **参数解释：** api-key的ID，在[创建API_KEY](CreateInferApiKey.xml)时即可在返回体中获取，也可通过[查询api-keys列表](ListInferApiKeys.xml)获取当前用户拥有的api-key，其中id字段即为api-key的ID。 **取值范围：** UUID格式。
	KeyId string `json:"key_id"`

	// **参数解释：** ModelArts错误码。 **取值范围：** 不涉及。
	ErrorCode string `json:"error_code"`

	// **参数解释：** 具体错误信息。 **取值范围：** 不涉及。
	ErrorMsg string `json:"error_msg"`
}

func (o ApiKeyFailureResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiKeyFailureResponse struct{}"
	}

	return strings.Join([]string{"ApiKeyFailureResponse", string(data)}, " ")
}
