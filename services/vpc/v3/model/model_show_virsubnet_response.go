package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVirsubnetResponse Response Object
type ShowVirsubnetResponse struct {

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId *string `json:"request_id,omitempty"`

	Virsubnet      *Virsubnet `json:"virsubnet,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowVirsubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVirsubnetResponse struct{}"
	}

	return strings.Join([]string{"ShowVirsubnetResponse", string(data)}, " ")
}
