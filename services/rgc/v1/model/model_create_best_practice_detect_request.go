package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBestPracticeDetectRequest Request Object
type CreateBestPracticeDetectRequest struct {
}

func (o CreateBestPracticeDetectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBestPracticeDetectRequest struct{}"
	}

	return strings.Join([]string{"CreateBestPracticeDetectRequest", string(data)}, " ")
}
