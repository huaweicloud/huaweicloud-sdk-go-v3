package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteMessageEmailConfigRequest struct {
}

func (o DeleteMessageEmailConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMessageEmailConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteMessageEmailConfigRequest", string(data)}, " ")
}
