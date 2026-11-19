package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostsUpgradeResponse Response Object
type ListHostsUpgradeResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2097152
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 服务器的升级状态列表 **取值范围**: 0-10241个HostUpgradeResponseInfo对象
	DataList       *[]HostUpgradeResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListHostsUpgradeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostsUpgradeResponse struct{}"
	}

	return strings.Join([]string{"ListHostsUpgradeResponse", string(data)}, " ")
}
