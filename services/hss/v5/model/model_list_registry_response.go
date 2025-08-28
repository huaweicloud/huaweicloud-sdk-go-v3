package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRegistryResponse Response Object
type ListRegistryResponse struct {

	// **参数解释**: 镜像仓列表 **取值范围**: 列表项数量0-100
	DataList *[]RegistryInfo `json:"data_list,omitempty"`

	// **参数解释**: 总数 **取值范围**: 0-100
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRegistryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegistryResponse struct{}"
	}

	return strings.Join([]string{"ListRegistryResponse", string(data)}, " ")
}
