package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRtcAbnormalEventResponse Response Object
type ListRtcAbnormalEventResponse struct {

	// 房间ID
	RoomId *string `json:"room_id,omitempty"`

	// 用户ID
	Uid *string `json:"uid,omitempty"`

	// 体验ID
	ExpId *string `json:"exp_id,omitempty"`

	// 异常信息列表。注意：此字段可能返回null，表示取不到有效值。
	AbnormalList *[]RtcCause `json:"abnormal_list,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRtcAbnormalEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRtcAbnormalEventResponse struct{}"
	}

	return strings.Join([]string{"ListRtcAbnormalEventResponse", string(data)}, " ")
}
