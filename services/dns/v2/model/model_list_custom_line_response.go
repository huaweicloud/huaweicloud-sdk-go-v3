package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCustomLineResponse struct {
	Lines *[]Line `json:"lines,omitempty" xml:"lines"`

	Metadata       *Metedata `json:"metadata,omitempty" xml:"metadata"`
	HttpStatusCode int       `json:"-"`
}

func (o ListCustomLineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomLineResponse struct{}"
	}

	return strings.Join([]string{"ListCustomLineResponse", string(data)}, " ")
}
