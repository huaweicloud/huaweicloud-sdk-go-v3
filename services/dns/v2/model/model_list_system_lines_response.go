package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSystemLinesResponse Response Object
type ListSystemLinesResponse struct {

	// **参数解释：** 线路列表。 **取值范围：** 不涉及。
	Lines *[]SystemLine `json:"lines,omitempty"`

	Metadata       *Metadata `json:"metadata,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSystemLinesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemLinesResponse struct{}"
	}

	return strings.Join([]string{"ListSystemLinesResponse", string(data)}, " ")
}
