package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobResponse Response Object
type ShowJobResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“Job”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	Spec           *JobSpec `json:"spec,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}
