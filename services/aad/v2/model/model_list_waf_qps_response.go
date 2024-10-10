package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWafQpsResponse Response Object
type ListWafQpsResponse struct {

	// curve
	Curve          *[]Point `json:"curve,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListWafQpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWafQpsResponse struct{}"
	}

	return strings.Join([]string{"ListWafQpsResponse", string(data)}, " ")
}
