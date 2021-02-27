package model

import (
	"encoding/json"

	"strings"
)

type RateOnPeriodReq struct {
	// 项目ID。 调用通过assume_role方式获取用户token接口获取客户Token后，参见如何将合作伙伴Token置换为客户Token的步骤2获取项目ID。
	ProjectId string `json:"project_id"`
	// 产品信息列表，询价时要询价产品的信息的列表，具体参见表1。
	ProductInfos []PeriodProductInfo `json:"product_infos"`
}

func (o RateOnPeriodReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RateOnPeriodReq struct{}"
	}

	return strings.Join([]string{"RateOnPeriodReq", string(data)}, " ")
}
