package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCatalogsResponse Response Object
type ListCatalogsResponse struct {
	Body           *[]Catalog `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListCatalogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCatalogsResponse struct{}"
	}

	return strings.Join([]string{"ListCatalogsResponse", string(data)}, " ")
}
