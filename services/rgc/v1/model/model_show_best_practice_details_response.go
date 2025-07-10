package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBestPracticeDetailsResponse Response Object
type ShowBestPracticeDetailsResponse struct {

	// 最佳实践检测结果
	Body           *[]BestPracticeCheckItemDetail `json:"body,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowBestPracticeDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBestPracticeDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowBestPracticeDetailsResponse", string(data)}, " ")
}
