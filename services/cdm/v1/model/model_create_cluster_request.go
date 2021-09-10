package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateClusterRequest struct {
	// 请求语言。

	XLanguage string `json:"X-Language"`

	Body *CdmCreateClusterReq `json:"body,omitempty"`
}

func (o CreateClusterRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateClusterRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterRequest", string(data)}, " ")
}
