package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SynchronizeInstanceListNewRequest Request Object
type SynchronizeInstanceListNewRequest struct {
}

func (o SynchronizeInstanceListNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SynchronizeInstanceListNewRequest struct{}"
	}

	return strings.Join([]string{"SynchronizeInstanceListNewRequest", string(data)}, " ")
}
