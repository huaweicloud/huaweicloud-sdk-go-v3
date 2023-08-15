package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseNamesResponse Response Object
type ListDatabaseNamesResponse struct {
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListDatabaseNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseNamesResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseNamesResponse", string(data)}, " ")
}
