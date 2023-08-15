package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectReleaseFilesResponse Response Object
type ShowProjectReleaseFilesResponse struct {
	Result         *StandardResponseResult `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowProjectReleaseFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectReleaseFilesResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectReleaseFilesResponse", string(data)}, " ")
}
