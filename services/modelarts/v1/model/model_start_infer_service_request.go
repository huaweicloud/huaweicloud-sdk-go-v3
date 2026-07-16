package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartInferServiceRequest Request Object
type StartInferServiceRequest struct {

	// **参数解释：** 服务ID，在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 服务ID。 **默认取值：** 不涉及。
	ServiceId string `json:"service_id"`
}

func (o StartInferServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInferServiceRequest struct{}"
	}

	return strings.Join([]string{"StartInferServiceRequest", string(data)}, " ")
}
