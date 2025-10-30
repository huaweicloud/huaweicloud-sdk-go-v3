package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateHostAssetValueRequestInfo struct {

	// **参数解释**： 资产重要性 **约束限制**： 不涉及 **取值范围**： - important：重要资产 - common：一般资产 - test：测试资产 **默认取值**： 不涉及
	AssetValue string `json:"asset_value"`

	// **参数解释**： 主机ID列表 **约束限制**： 不涉及 **取值范围**： 列表大小为1-200条 **默认取值**： 不涉及
	HostIdList []string `json:"host_id_list"`
}

func (o AssociateHostAssetValueRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateHostAssetValueRequestInfo struct{}"
	}

	return strings.Join([]string{"AssociateHostAssetValueRequestInfo", string(data)}, " ")
}
