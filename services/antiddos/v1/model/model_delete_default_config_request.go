package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDefaultConfigRequest Request Object
type DeleteDefaultConfigRequest struct {
}

func (o DeleteDefaultConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDefaultConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteDefaultConfigRequest", string(data)}, " ")
}
