package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProjectInfoV4Response struct {
	Project        *GetProjectInfoV4ResultProject `json:"project,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowProjectInfoV4Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectInfoV4Response struct{}"
	}

	return strings.Join([]string{"ShowProjectInfoV4Response", string(data)}, " ")
}
