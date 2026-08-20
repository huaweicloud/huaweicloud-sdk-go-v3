package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainLocationStatsResponse Response Object
type ShowDomainLocationStatsResponse struct {

	// **参数解释：** 数据分组方式 **取值范围：** - domain：按域名分组 - country：按国际&地区分组 - province：按省份分组 - isp：按运营商分组
	GroupBy *string `json:"group_by,omitempty"`

	// **参数解释：** 按指定的分组方式组织的数据 **取值范围：** 不涉及
	Result         map[string]interface{} `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowDomainLocationStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainLocationStatsResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainLocationStatsResponse", string(data)}, " ")
}
