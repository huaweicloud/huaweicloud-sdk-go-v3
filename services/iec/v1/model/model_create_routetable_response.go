package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoutetableResponse Response Object
type CreateRoutetableResponse struct {
	Routetable     *Routetable `json:"routetable,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateRoutetableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoutetableResponse struct{}"
	}

	return strings.Join([]string{"CreateRoutetableResponse", string(data)}, " ")
}
