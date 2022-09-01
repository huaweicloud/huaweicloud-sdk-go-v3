package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcAbnormalEventDimensionResponse struct {

	// 异常体验列表
	Dimensions *[]AbnormalEventDimensionValue `json:"dimensions,omitempty" xml:"dimensions"`

	XRequestId     *string `json:"X-request-id,omitempty" xml:"X-request-id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcAbnormalEventDimensionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcAbnormalEventDimensionResponse struct{}"
	}

	return strings.Join([]string{"ListRtcAbnormalEventDimensionResponse", string(data)}, " ")
}
