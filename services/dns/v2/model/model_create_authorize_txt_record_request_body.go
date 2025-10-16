package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateAuthorizeTxtRecordRequestBody struct {

	// **参数解释：** 待创建的子域名。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ZoneName string `json:"zone_name"`
}

func (o CreateAuthorizeTxtRecordRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuthorizeTxtRecordRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAuthorizeTxtRecordRequestBody", string(data)}, " ")
}
