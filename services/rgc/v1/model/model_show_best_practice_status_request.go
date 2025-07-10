package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBestPracticeStatusRequest Request Object
type ShowBestPracticeStatusRequest struct {
}

func (o ShowBestPracticeStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBestPracticeStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowBestPracticeStatusRequest", string(data)}, " ")
}
