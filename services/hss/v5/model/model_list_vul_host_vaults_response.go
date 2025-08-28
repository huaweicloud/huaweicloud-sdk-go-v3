package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHostVaultsResponse Response Object
type ListVulHostVaultsResponse struct {

	// **参数解释**: 总数 **取值范围**: 取值0-10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 主机存储库信息列表 **取值范围**: 不涉及
	DataList       *[]HostVaultInfo `json:"data_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListVulHostVaultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHostVaultsResponse struct{}"
	}

	return strings.Join([]string{"ListVulHostVaultsResponse", string(data)}, " ")
}
