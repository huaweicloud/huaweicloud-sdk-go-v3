package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDefaultConfigRequest struct {
}

func (o DeleteDefaultConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDefaultConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteDefaultConfigRequest", string(data)}, " ")
}
