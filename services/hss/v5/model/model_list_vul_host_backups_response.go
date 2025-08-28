package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulHostBackupsResponse Response Object
type ListVulHostBackupsResponse struct {

	// **参数解释**: 总数 **取值范围**: 取值0-10000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 备份列表 **取值范围**: 不涉及
	DataList       *[]VulHostBackupsResponseInfoDataList `json:"data_list,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListVulHostBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulHostBackupsResponse struct{}"
	}

	return strings.Join([]string{"ListVulHostBackupsResponse", string(data)}, " ")
}
