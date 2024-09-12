package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDigitalHumanVideoResponse Response Object
type ListDigitalHumanVideoResponse struct {

	// **参数解释**： 视频制作任务总数。
	Count *int32 `json:"count,omitempty"`

	// 视频制作任务列表。
	Jobs *[]DigitalHumanVideo `json:"jobs,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDigitalHumanVideoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDigitalHumanVideoResponse struct{}"
	}

	return strings.Join([]string{"ListDigitalHumanVideoResponse", string(data)}, " ")
}
