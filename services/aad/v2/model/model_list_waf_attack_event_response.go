package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWafAttackEventResponse Response Object
type ListWafAttackEventResponse struct {

	// total
	Total *int32 `json:"total,omitempty"`

	// list
	List           *[]ListWafAttackEventlist `json:"list,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListWafAttackEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWafAttackEventResponse struct{}"
	}

	return strings.Join([]string{"ListWafAttackEventResponse", string(data)}, " ")
}
