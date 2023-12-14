package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumeSimultaneousInterpretationResponse Response Object
type ResumeSimultaneousInterpretationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResumeSimultaneousInterpretationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeSimultaneousInterpretationResponse struct{}"
	}

	return strings.Join([]string{"ResumeSimultaneousInterpretationResponse", string(data)}, " ")
}
