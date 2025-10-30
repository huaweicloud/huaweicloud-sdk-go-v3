package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulRepoImageResponse Response Object
type ListVulRepoImageResponse struct {

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483547
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**: 仓库镜像列表 **取值范围**: 最小值0，最大值10000
	DataList       *[]VulRepoImageInfo `json:"data_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListVulRepoImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulRepoImageResponse struct{}"
	}

	return strings.Join([]string{"ListVulRepoImageResponse", string(data)}, " ")
}
