package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAnalyzerResponse Response Object
type DeleteAnalyzerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAnalyzerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAnalyzerResponse struct{}"
	}

	return strings.Join([]string{"DeleteAnalyzerResponse", string(data)}, " ")
}
