package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAnalyzerResponse Response Object
type UpdateAnalyzerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateAnalyzerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAnalyzerResponse struct{}"
	}

	return strings.Join([]string{"UpdateAnalyzerResponse", string(data)}, " ")
}
