package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListScriptResultsResponse struct {
	Status *string `json:"status,omitempty"`

	Result         *[]Result `json:"result,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListScriptResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScriptResultsResponse struct{}"
	}

	return strings.Join([]string{"ListScriptResultsResponse", string(data)}, " ")
}
