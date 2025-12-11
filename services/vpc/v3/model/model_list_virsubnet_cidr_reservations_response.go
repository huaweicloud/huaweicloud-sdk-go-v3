package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVirsubnetCidrReservationsResponse Response Object
type ListVirsubnetCidrReservationsResponse struct {

	// **参数解释**： 查询子网预留网段列表的响应体。 **取值范围**： 不涉及。
	VirsubnetCidrReservations *[]VirsubnetCidrReservation `json:"virsubnet_cidr_reservations,omitempty"`

	// **参数解释**： 请求ID。 **取值范围**： 不涉及。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVirsubnetCidrReservationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVirsubnetCidrReservationsResponse struct{}"
	}

	return strings.Join([]string{"ListVirsubnetCidrReservationsResponse", string(data)}, " ")
}
