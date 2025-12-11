package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstancePeriodRequestProductInfos struct {

	// 产品ID
	ProductId string `json:"product_id"`

	// 服务类型： - hws.service.type.dbss: 数据库审计
	CloudServiceType string `json:"cloud_service_type"`

	// 资源类型: - hws.resource.type.dbss：数据库审计
	ResourceType string `json:"resource_type"`

	// 资源规格： - dbss.bypassaudit.low：基础版 - dbss.bypassaudit.medium：高级版 - dbss.bypassaudit.high：专业版
	ResourceSpecCode string `json:"resource_spec_code"`

	// 产品规格描述。json字符串格式 ：{\"specDesc\":{\"zh-cn\":{\"key1\":\"value1\"},\"en-us\":{\"key1\":\"value1\"}}}
	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`
}

func (o CreateInstancePeriodRequestProductInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstancePeriodRequestProductInfos struct{}"
	}

	return strings.Join([]string{"CreateInstancePeriodRequestProductInfos", string(data)}, " ")
}
