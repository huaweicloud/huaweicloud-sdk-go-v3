package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobResponse Response Object
type ShowJobResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *JobKindObj `json:"kind,omitempty"`

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
