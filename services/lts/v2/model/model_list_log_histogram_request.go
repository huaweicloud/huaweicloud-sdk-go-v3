package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLogHistogramRequest struct {
	Body *QueryLogKeyWordCountRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ListLogHistogramRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogHistogramRequest struct{}"
	}

	return strings.Join([]string{"ListLogHistogramRequest", string(data)}, " ")
}
