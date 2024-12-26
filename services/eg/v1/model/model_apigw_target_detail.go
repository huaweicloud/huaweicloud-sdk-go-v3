package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApigwTargetDetail struct {

	// 目标url
	Url string `json:"url"`

	InvocationHttpParameters *InvocationHttpParameters `json:"invocation_http_parameters,omitempty"`
}

func (o ApigwTargetDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApigwTargetDetail struct{}"
	}

	return strings.Join([]string{"ApigwTargetDetail", string(data)}, " ")
}
