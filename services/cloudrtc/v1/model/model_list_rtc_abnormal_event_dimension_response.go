package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcAbnormalEventDimensionResponse Response Object
type ListRtcAbnormalEventDimensionResponse struct {

	// 异常体验列表
	Dimensions *[]AbnormalEventDimensionValue `json:"dimensions,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcAbnormalEventDimensionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcAbnormalEventDimensionResponse struct{}"
	}

	return strings.Join([]string{"ListRtcAbnormalEventDimensionResponse", string(data)}, " ")
}
