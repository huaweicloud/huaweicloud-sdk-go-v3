package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRtcClientQosDetailsResponse struct {

	// 房间ID
	RoomId *string `json:"room_id,omitempty" xml:"room_id"`

	// QoS质量数据
	Data *[]QosQualityData `json:"data,omitempty" xml:"data"`

	XRequestId     *string `json:"X-request-id,omitempty" xml:"X-request-id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcClientQosDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcClientQosDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListRtcClientQosDetailsResponse", string(data)}, " ")
}
