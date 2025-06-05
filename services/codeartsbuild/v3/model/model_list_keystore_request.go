package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKeystoreRequest Request Object
type ListKeystoreRequest struct {
}

func (o ListKeystoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeystoreRequest struct{}"
	}

	return strings.Join([]string{"ListKeystoreRequest", string(data)}, " ")
}
