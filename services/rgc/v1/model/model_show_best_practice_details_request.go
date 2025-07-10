package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBestPracticeDetailsRequest Request Object
type ShowBestPracticeDetailsRequest struct {
}

func (o ShowBestPracticeDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBestPracticeDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowBestPracticeDetailsRequest", string(data)}, " ")
}
