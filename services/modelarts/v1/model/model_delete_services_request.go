package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteServicesRequest 提供待删除的services id列表。
type DeleteServicesRequest struct {

	// **参数解释：** 待删除的servicesid列表。服务ID在[创建服务](CreateInferService.xml)时即可在返回体中获取，也可通过[查询服务列表](ListInferServices.xml)获取当前用户拥有的服务，其中service_id字段即为服务ID。 **约束限制：** 不涉及。 **取值范围：** 服务ID。 **默认取值：** 不涉及。
	ServiceIds []string `json:"service_ids"`
}

func (o DeleteServicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteServicesRequest struct{}"
	}

	return strings.Join([]string{"DeleteServicesRequest", string(data)}, " ")
}
