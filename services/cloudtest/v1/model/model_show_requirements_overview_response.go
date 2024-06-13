package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRequirementsOverviewResponse Response Object
type ShowRequirementsOverviewResponse struct {

	// success|error;
	Status *string `json:"status,omitempty"`

	Result *ResultValueRequirementsOverviewVo `json:"result,omitempty"`

	Error          *ApiError `json:"error,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowRequirementsOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRequirementsOverviewResponse struct{}"
	}

	return strings.Join([]string{"ShowRequirementsOverviewResponse", string(data)}, " ")
}
