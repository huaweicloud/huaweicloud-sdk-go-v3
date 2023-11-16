package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectAsyncStatusLogInfoRequest Request Object
type ShowProjectAsyncStatusLogInfoRequest struct {
}

func (o ShowProjectAsyncStatusLogInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectAsyncStatusLogInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectAsyncStatusLogInfoRequest", string(data)}, " ")
}
