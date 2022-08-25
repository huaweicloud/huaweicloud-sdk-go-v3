package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FilterFactor struct {

	// 维度分组条件。CLOUD_SERVICE_TYPE：产品类型ASSOCIATED_ACCOUNT：关联账号REGION_CODE：区域RES_SPEC_CODE：规格编码USAGE_TYPE：使用量类型ENTERPRISE_PROJECT_ID：企业项目RESOURCE_ID：资源CHARGING_MODE：计费模式BILL_TYPE：账单类型RESOURCE_TYPE：产品AZ_CODE：可用区BE_ID：运营实体（beid）PAYER_ACCOUNT_ID：交易账号RESOURCE_TAG：成本标签COST_UNIT：成本单元
	Key string `json:"key"`

	// 过滤器值
	Value []string `json:"value"`
}

func (o FilterFactor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FilterFactor struct{}"
	}

	return strings.Join([]string{"FilterFactor", string(data)}, " ")
}
