package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRoutetableResponse Response Object
type UpdateRoutetableResponse struct {
	Routetable     *Routetable `json:"routetable,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o UpdateRoutetableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRoutetableResponse struct{}"
	}

	return strings.Join([]string{"UpdateRoutetableResponse", string(data)}, " ")
}
