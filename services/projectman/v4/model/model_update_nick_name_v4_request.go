package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateNickNameV4Request struct {
	Body *UpdateUserNickNameRequestV4 `json:"body,omitempty"`
}

func (o UpdateNickNameV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNickNameV4Request struct{}"
	}

	return strings.Join([]string{"UpdateNickNameV4Request", string(data)}, " ")
}
