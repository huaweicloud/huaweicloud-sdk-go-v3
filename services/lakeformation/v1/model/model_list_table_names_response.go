package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTableNamesResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTableNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableNamesResponse struct{}"
	}

	return strings.Join([]string{"ListTableNamesResponse", string(data)}, " ")
}
