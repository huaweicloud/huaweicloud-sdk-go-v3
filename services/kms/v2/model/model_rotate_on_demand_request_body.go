package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RotateOnDemandRequestBody struct {

	// **参数解释：** 密钥ID **约束限制：** - 36字节ID - 满足正则匹配“^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$” **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyId string `json:"key_id"`
}

func (o RotateOnDemandRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateOnDemandRequestBody struct{}"
	}

	return strings.Join([]string{"RotateOnDemandRequestBody", string(data)}, " ")
}
