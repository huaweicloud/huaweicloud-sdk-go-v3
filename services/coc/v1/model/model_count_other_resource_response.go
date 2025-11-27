package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountOtherResourceResponse Response Object
type CountOtherResourceResponse struct {

	// **参数解释：** IDC线下资源总数。 **取值范围：** 不涉及。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CountOtherResourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountOtherResourceResponse struct{}"
	}

	return strings.Join([]string{"CountOtherResourceResponse", string(data)}, " ")
}
