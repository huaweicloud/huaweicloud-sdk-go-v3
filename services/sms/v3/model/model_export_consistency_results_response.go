package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportConsistencyResultsResponse Response Object
type ExportConsistencyResultsResponse struct {

	// OK
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportConsistencyResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportConsistencyResultsResponse struct{}"
	}

	return strings.Join([]string{"ExportConsistencyResultsResponse", string(data)}, " ")
}
