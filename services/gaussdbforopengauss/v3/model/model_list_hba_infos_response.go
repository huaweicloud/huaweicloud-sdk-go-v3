package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHbaInfosResponse Response Object
type ListHbaInfosResponse struct {

	// **参数解释**: 客户端接入配置列表。
	HbaConfs *[]HbaConfResult `json:"hba_confs,omitempty"`

	// **参数解释**: hba配置总数。 **取值范围**: [0, 2^31-1]，取决于实际查询大小。
	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListHbaInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHbaInfosResponse struct{}"
	}

	return strings.Join([]string{"ListHbaInfosResponse", string(data)}, " ")
}
