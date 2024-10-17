package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstancePeriodRequestProductInfos struct {

	// 产品ID
	ProductId string `json:"product_id"`

	// 服务类型： - hws.service.type.dbss
	CloudServiceType string `json:"cloud_service_type"`

	// 资源类型: - hws.resource.type.dbss
	ResourceType string `json:"resource_type"`

	// 资源规格： - dbss.bypassaudit.low - dbss.bypassaudit.medium - dbss.bypassaudit.high
	ResourceSpecCode string `json:"resource_spec_code"`

	// 产品规格描述。json字符串格式 ：{\"specDesc\":{\"zh-cn\":{\"key1\":\"value1\"},\"en-us\":{\"key1\":\"value1\"}}}。key和value为用户自定义希望展示的产品描述信息。
	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`
}

func (o CreateInstancePeriodRequestProductInfos) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstancePeriodRequestProductInfos struct{}"
	}

	return strings.Join([]string{"CreateInstancePeriodRequestProductInfos", string(data)}, " ")
}
