package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFlinkSavepointResponse Response Object
type CreateFlinkSavepointResponse struct {

	// **参数解释**：Savepoint唯一标识。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。
	SavepointId    *string `json:"savepoint_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateFlinkSavepointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFlinkSavepointResponse struct{}"
	}

	return strings.Join([]string{"CreateFlinkSavepointResponse", string(data)}, " ")
}
