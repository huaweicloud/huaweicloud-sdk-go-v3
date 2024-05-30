package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DimensionGroup struct {

	// 分组条件。 产品类型：CLOUD_SERVICE_TYPE企业项目：ENTERPRISE_PROJECT_ID使用量类型：USAGE_TYPE产品：RESOURCE_TYPE可用区：AZ_CODE账单类型：BILL_TYPE关联账号：ASSOCIATED_ACCOUNT规格编码：RES_SPEC_CODE运营实体：BE_ID区域：REGION_CODE计费模式：CHARGING_MODE交易账号：PAYER_ACCOUNT_ID资源tag：RESOURCE_TAG资源id：RESOURCE_ID成本单元：COST_UNIT分拆项：SPLIT_ITEM_NAME
	Key *string `json:"key,omitempty"`

	// 维度值。
	Value *string `json:"value,omitempty"`
}

func (o DimensionGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionGroup struct{}"
	}

	return strings.Join([]string{"DimensionGroup", string(data)}, " ")
}
