package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogHistogramRequest Request Object
type ListLogHistogramRequest struct {
	Body *QueryLogKeyWordCountRequestBody `json:"body,omitempty"`
}

func (o ListLogHistogramRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogHistogramRequest struct{}"
	}

	return strings.Join([]string{"ListLogHistogramRequest", string(data)}, " ")
}
