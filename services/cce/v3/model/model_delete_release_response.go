package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteReleaseResponse Response Object
type DeleteReleaseResponse struct {

	// **参数解释**： 删除模板实例的返回结果。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteReleaseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteReleaseResponse struct{}"
	}

	return strings.Join([]string{"DeleteReleaseResponse", string(data)}, " ")
}
