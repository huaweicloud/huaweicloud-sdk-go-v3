package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQuotaStatusRequest Request Object
type ListQuotaStatusRequest struct {

	// 创建云堡垒机所在的可用区，需要指定可用分区名称。 可参考[地区和终端节点](https://developer.huaweicloud.com/endpoint)获取
	AvailabilityZone string `json:"availability_zone"`

	// 待创建云堡垒机规格ID，例如： - cbh.basic.10  10资产标准版 - cbh.enhance.10  10资产专业版 已上线的规格详情请参见《云堡垒机常见问题》的购买，[云堡垒机实例有哪些规格](https://support.huaweicloud.com/cbh_faq/cbh_03_0025.html)章节。
	ResourceSpecCode string `json:"resource_spec_code"`
}

func (o ListQuotaStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQuotaStatusRequest struct{}"
	}

	return strings.Join([]string{"ListQuotaStatusRequest", string(data)}, " ")
}
