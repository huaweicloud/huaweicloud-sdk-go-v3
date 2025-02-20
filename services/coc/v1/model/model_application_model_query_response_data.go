package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicationModelQueryResponseData struct {
	SubApplications *[]ApplicationInfo `json:"sub_applications,omitempty"`

	Components *[]ComponentInfo `json:"components,omitempty"`

	Groups *[]GroupInfo `json:"groups,omitempty"`
}

func (o ApplicationModelQueryResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationModelQueryResponseData struct{}"
	}

	return strings.Join([]string{"ApplicationModelQueryResponseData", string(data)}, " ")
}
