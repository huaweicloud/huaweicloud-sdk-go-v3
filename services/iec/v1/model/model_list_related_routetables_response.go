package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRelatedRoutetablesResponse struct {
	Routetable     *Routetable `json:"routetable,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListRelatedRoutetablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelatedRoutetablesResponse struct{}"
	}

	return strings.Join([]string{"ListRelatedRoutetablesResponse", string(data)}, " ")
}
