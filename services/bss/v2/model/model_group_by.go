package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GroupBy struct {

	// tag：按照成本标签过滤cost_unit：按照成本单元过滤dimension：默认取值
	Type string `json:"type"`

	// 如果type为tag，此处取值为tag的key。如果type为cost_unit，此处取值为cost_unit的key。如果type为dimension，此处取值如下：CLOUD_SERVICE_TYPE 产品类型RESOURCE_TYPE 产品ASSOCIATED_ACCOUNT 关联账号（企业主名下的所有企业子账号）REGION_CODE 区域AZ_CODE 可用区ENTERPRISE_PROJECT_ID 企业项目RES_SPEC_CODE 产品规格CHARGING_MODE 计费模式USAGE_TYPE 使用量类型BILL_TYPE 账单大类（billtype，前台需转换billdetailtype）BE_ID 运营实体（beid）PAYER_ACCOUNT_ID 支付账号RESOURCE_ID 资源
	Key string `json:"key"`
}

func (o GroupBy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupBy struct{}"
	}

	return strings.Join([]string{"GroupBy", string(data)}, " ")
}
