package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProfileAssociationsResponse Response Object
type ListProfileAssociationsResponse struct {
	Associations   *[]AssociationDto `json:"associations,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListProfileAssociationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProfileAssociationsResponse struct{}"
	}

	return strings.Join([]string{"ListProfileAssociationsResponse", string(data)}, " ")
}
