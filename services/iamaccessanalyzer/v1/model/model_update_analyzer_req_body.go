package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAnalyzerReqBody struct {
	Configuration *AnalyzerConfiguration `json:"configuration,omitempty"`
}

func (o UpdateAnalyzerReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnalyzerReqBody struct{}"
	}

	return strings.Join([]string{"UpdateAnalyzerReqBody", string(data)}, " ")
}
