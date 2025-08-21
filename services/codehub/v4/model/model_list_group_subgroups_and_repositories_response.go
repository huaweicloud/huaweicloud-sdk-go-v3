package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupSubgroupsAndRepositoriesResponse Response Object
type ListGroupSubgroupsAndRepositoriesResponse struct {

	// **参数解释：** 项目下代码组和仓库列表。
	Body           *[]SubgroupAndProjectBaseDto `json:"body,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListGroupSubgroupsAndRepositoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupSubgroupsAndRepositoriesResponse struct{}"
	}

	return strings.Join([]string{"ListGroupSubgroupsAndRepositoriesResponse", string(data)}, " ")
}
