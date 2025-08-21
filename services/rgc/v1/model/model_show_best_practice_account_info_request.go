package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBestPracticeAccountInfoRequest Request Object
type ShowBestPracticeAccountInfoRequest struct {
}

func (o ShowBestPracticeAccountInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBestPracticeAccountInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowBestPracticeAccountInfoRequest", string(data)}, " ")
}
