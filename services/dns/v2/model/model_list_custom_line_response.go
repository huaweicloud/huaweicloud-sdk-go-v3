package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCustomLineResponse struct {
	Lines *[]Line `json:"lines,omitempty"`

	Metadata       *Metedata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCustomLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomLineResponse struct{}"
	}

	return strings.Join([]string{"ListCustomLineResponse", string(data)}, " ")
}
