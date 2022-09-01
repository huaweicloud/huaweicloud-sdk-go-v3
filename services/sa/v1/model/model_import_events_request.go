package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportEventsRequest struct {
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	Body *Events `json:"body,omitempty" xml:"body"`
}

func (o ImportEventsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportEventsRequest struct{}"
	}

	return strings.Join([]string{"ImportEventsRequest", string(data)}, " ")
}
