package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRoutetableResponse struct {
	Routetable     *Routetable `json:"routetable,omitempty" xml:"routetable"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowRoutetableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRoutetableResponse struct{}"
	}

	return strings.Join([]string{"ShowRoutetableResponse", string(data)}, " ")
}
