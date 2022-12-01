package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateRequestRequest struct {

	// the number of samples
	Topn *int32 `json:"topn,omitempty"`

	// the scenario of code content
	Scenario *string `json:"scenario,omitempty"`

	// if `resubmit` is true, the de-duplication will be ignored
	Resubmit *bool `json:"resubmit,omitempty"`

	// choose the model
	ModelId *string `json:"model_id,omitempty"`

	Body *PropertiesSchema `json:"body,omitempty"`
}

func (o CreateRequestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRequestRequest struct{}"
	}

	return strings.Join([]string{"CreateRequestRequest", string(data)}, " ")
}
