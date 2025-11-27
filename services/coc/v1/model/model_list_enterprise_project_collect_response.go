package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnterpriseProjectCollectResponse Response Object
type ListEnterpriseProjectCollectResponse struct {

	// **参数解释：** 收藏对象列表。 **取值范围：** 列表大小0~500。
	Data           *[]EnterpriseProjectCollectQueryResponseData `json:"data,omitempty"`
	HttpStatusCode int                                          `json:"-"`
}

func (o ListEnterpriseProjectCollectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnterpriseProjectCollectResponse struct{}"
	}

	return strings.Join([]string{"ListEnterpriseProjectCollectResponse", string(data)}, " ")
}
