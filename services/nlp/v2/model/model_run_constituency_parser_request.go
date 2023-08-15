package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunConstituencyParserRequest Request Object
type RunConstituencyParserRequest struct {
	Body *ConstituencyParserReq `json:"body,omitempty"`
}

func (o RunConstituencyParserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunConstituencyParserRequest struct{}"
	}

	return strings.Join([]string{"RunConstituencyParserRequest", string(data)}, " ")
}
