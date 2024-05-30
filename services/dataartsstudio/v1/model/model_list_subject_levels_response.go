package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubjectLevelsResponse Response Object
type ListSubjectLevelsResponse struct {
	Data           *ListSubjectLevelsResultData `json:"data,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListSubjectLevelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubjectLevelsResponse struct{}"
	}

	return strings.Join([]string{"ListSubjectLevelsResponse", string(data)}, " ")
}
