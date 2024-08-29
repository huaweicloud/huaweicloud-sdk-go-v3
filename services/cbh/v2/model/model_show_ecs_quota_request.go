package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEcsQuotaRequest Request Object
type ShowEcsQuotaRequest struct {

	// 可用分区名称。  可参考接口\"获取服务可用区\"获取
	AvailabilityZone string `json:"availability_zone"`

	// 待创建云堡垒机规格ID，例如： - cbh.basic.10  10资产标准版 - cbh.enhance.10  10资产专业版  可参考接口\"查询云堡垒机规格信息\"获取
	ResourceSpecCode string `json:"resource_spec_code"`
}

func (o ShowEcsQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEcsQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowEcsQuotaRequest", string(data)}, " ")
}
