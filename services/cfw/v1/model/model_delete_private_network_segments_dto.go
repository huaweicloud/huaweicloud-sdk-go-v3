package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeletePrivateNetworkSegmentsDto struct {

	// **参数解释**： 删除的私网网段ID列表，可以通过调用[获取私网网段的信息]获得，通过返回值中的data.private_network_list.conf_id获得 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ConfIds []string `json:"conf_ids"`
}

func (o DeletePrivateNetworkSegmentsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateNetworkSegmentsDto struct{}"
	}

	return strings.Join([]string{"DeletePrivateNetworkSegmentsDto", string(data)}, " ")
}
