package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobStatusResponse Response Object
type ShowJobStatusResponse struct {

	// 作业运行信息，详见submissions参数说明。
	Submissions    *[]Submission `json:"submissions,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowJobStatusResponse", string(data)}, " ")
}
