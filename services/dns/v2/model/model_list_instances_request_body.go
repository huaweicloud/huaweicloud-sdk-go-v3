package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListInstancesRequestBody struct {

	// **参数解释：** DNS批量查询接口支持查询的维度。 **约束限制：** 不涉及。 **取值范围：** - dns_public_zone_id：公网域名ID - dns_public_recordset_id：公网记录集ID，需与dns_public_zone_id组合使用 - dns_private_zone_id：内网域名ID - vpc_id：VPC ID，需与dns_private_zone_id组合使用 **默认取值：** 不涉及。
	DimName []string `json:"dim_name"`

	// DNS上报的资源ID列表。
	Ids [][]string `json:"ids"`
}

func (o ListInstancesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequestBody struct{}"
	}

	return strings.Join([]string{"ListInstancesRequestBody", string(data)}, " ")
}
