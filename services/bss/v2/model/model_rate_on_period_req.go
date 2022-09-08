package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RateOnPeriodReq struct {

	// 项目ID。  说明： 获取方法： 步骤1：调用IAM获取委托Token接口，获取客户Token。步骤2：参见如何将合作伙伴Token置换为客户Token的步骤2，获取项目ID。IAM子用户调用此接口，需要IAM主账号授权，具体请参考创建用户组并授权。
	ProjectId string `json:"project_id"`

	// 产品信息列表，询价时要询价产品的信息的列表，具体参见表1。
	ProductInfos []PeriodProductInfo `json:"product_infos"`
}

func (o RateOnPeriodReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RateOnPeriodReq struct{}"
	}

	return strings.Join([]string{"RateOnPeriodReq", string(data)}, " ")
}
