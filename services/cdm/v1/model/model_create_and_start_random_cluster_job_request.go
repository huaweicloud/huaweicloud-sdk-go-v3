package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateAndStartRandomClusterJobRequest struct {
	// 请求语言。

	XLanguage string `json:"X-Language"`

	Body *CdmRandomCreateAndStartJobJsonReq `json:"body,omitempty"`
}

func (o CreateAndStartRandomClusterJobRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAndStartRandomClusterJobRequest struct{}"
	}

	return strings.Join([]string{"CreateAndStartRandomClusterJobRequest", string(data)}, " ")
}
