package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupabaseApplicationFlavorsResponse Response Object
type ListSupabaseApplicationFlavorsResponse struct {
	Flavors        *[]SupabaseSpecFlavorInfo `json:"flavors,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListSupabaseApplicationFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupabaseApplicationFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListSupabaseApplicationFlavorsResponse", string(data)}, " ")
}
