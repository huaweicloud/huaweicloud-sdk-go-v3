package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectSubgroupsAndRepositoriesResponse Response Object
type ListProjectSubgroupsAndRepositoriesResponse struct {

	// **参数解释：** 项目下代码组和仓库列表。
	Body           *[]SubgroupAndProjectBaseDto `json:"body,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListProjectSubgroupsAndRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectSubgroupsAndRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListProjectSubgroupsAndRepositoriesResponse", string(data)}, " ")
}
