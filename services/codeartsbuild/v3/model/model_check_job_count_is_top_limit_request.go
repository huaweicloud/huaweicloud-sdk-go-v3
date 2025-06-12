package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckJobCountIsTopLimitRequest Request Object
type CheckJobCountIsTopLimitRequest struct {
}

func (o CheckJobCountIsTopLimitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckJobCountIsTopLimitRequest struct{}"
	}

	return strings.Join([]string{"CheckJobCountIsTopLimitRequest", string(data)}, " ")
}
