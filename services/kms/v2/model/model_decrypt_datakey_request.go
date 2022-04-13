package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DecryptDatakeyRequest struct {
	Body *DecryptDatakeyRequestBody `json:"body,omitempty"`
}

func (o DecryptDatakeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDatakeyRequest struct{}"
	}

	return strings.Join([]string{"DecryptDatakeyRequest", string(data)}, " ")
}
