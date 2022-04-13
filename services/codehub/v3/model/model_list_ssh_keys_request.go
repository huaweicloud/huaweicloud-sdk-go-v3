package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSshKeysRequest struct {
}

func (o ListSshKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSshKeysRequest struct{}"
	}

	return strings.Join([]string{"ListSshKeysRequest", string(data)}, " ")
}
