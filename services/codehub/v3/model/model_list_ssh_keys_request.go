package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSshKeysRequest Request Object
type ListSshKeysRequest struct {
}

func (o ListSshKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSshKeysRequest struct{}"
	}

	return strings.Join([]string{"ListSshKeysRequest", string(data)}, " ")
}
