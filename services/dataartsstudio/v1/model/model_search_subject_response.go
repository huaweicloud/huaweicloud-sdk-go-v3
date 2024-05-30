package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSubjectResponse Response Object
type SearchSubjectResponse struct {
	Data           *SearchSubjectResultData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o SearchSubjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSubjectResponse struct{}"
	}

	return strings.Join([]string{"SearchSubjectResponse", string(data)}, " ")
}
