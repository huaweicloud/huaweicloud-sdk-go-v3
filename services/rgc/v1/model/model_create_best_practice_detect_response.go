package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBestPracticeDetectResponse Response Object
type CreateBestPracticeDetectResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateBestPracticeDetectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBestPracticeDetectResponse struct{}"
	}

	return strings.Join([]string{"CreateBestPracticeDetectResponse", string(data)}, " ")
}
