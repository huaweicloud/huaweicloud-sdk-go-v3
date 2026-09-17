package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEpsResponse Response Object
type ListEpsResponse struct {

	// **参数解释**： 集群及关联的企业项目信息。 **取值范围**： 不涉及。
	Resources *[]SysTagResp `json:"resources,omitempty"`

	// **参数解释**： 分页总条数。 **取值范围**： 大于等于0
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEpsResponse struct{}"
	}

	return strings.Join([]string{"ListEpsResponse", string(data)}, " ")
}
