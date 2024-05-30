package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSubjectNewResponse Response Object
type SearchSubjectNewResponse struct {
	Data           *SearchSubjectNewResultData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o SearchSubjectNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSubjectNewResponse struct{}"
	}

	return strings.Join([]string{"SearchSubjectNewResponse", string(data)}, " ")
}
