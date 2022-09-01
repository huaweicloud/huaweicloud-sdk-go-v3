package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RunDependencyParserRequest struct {
	Body *DependencyParserRequest `json:"body,omitempty" xml:"body"`
}

func (o RunDependencyParserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunDependencyParserRequest struct{}"
	}

	return strings.Join([]string{"RunDependencyParserRequest", string(data)}, " ")
}
