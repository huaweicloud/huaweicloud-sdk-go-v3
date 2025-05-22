package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRedistributionSchemaResponse Response Object
type ListRedistributionSchemaResponse struct {

	// **参数解释**： 错误码。 **取值范围**： 请求正常时为固定值 DWS.0000。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**： 错误信息。 **取值范围**： 不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**： 待重分布的schema列表。 **取值范围**： 不涉及。
	Schemas *[]RedisSchema `json:"schemas,omitempty"`

	// **参数解释**： 数据总条数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListRedistributionSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRedistributionSchemaResponse struct{}"
	}

	return strings.Join([]string{"ListRedistributionSchemaResponse", string(data)}, " ")
}
