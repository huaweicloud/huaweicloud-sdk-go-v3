package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePeriodElasticResourcePoolSpecChangeOrderRequestBody struct {

	// 弹性资源池名称，名称只能包含数字、小写英文字母和下划线，但不能是纯数字，且不能以下划线开头。长度限制：1~128个字符。
	ElasticResourcePoolName string `json:"elastic_resource_pool_name"`

	// CBC订单跳转DLI console链接
	CloudServiceConsoleUrl *string `json:"cloud_service_console_url,omitempty"`

	// 优惠信息
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 包周期目标CU大小
	TargetCu int32 `json:"target_cu"`
}

func (o CreatePeriodElasticResourcePoolSpecChangeOrderRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePeriodElasticResourcePoolSpecChangeOrderRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePeriodElasticResourcePoolSpecChangeOrderRequestBody", string(data)}, " ")
}
