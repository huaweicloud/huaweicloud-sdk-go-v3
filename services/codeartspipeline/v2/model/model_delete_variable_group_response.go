package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVariableGroupResponse Response Object
type DeleteVariableGroupResponse struct {

	// **参数解释**： 操作是否成功。 **取值范围**： - true：操作成功。 - false：操作失败。
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteVariableGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVariableGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteVariableGroupResponse", string(data)}, " ")
}
