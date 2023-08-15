package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowReleaseProjectFilesResponse Response Object
type ShowReleaseProjectFilesResponse struct {
	Result         *StandardResponseResult `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowReleaseProjectFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReleaseProjectFilesResponse struct{}"
	}

	return strings.Join([]string{"ShowReleaseProjectFilesResponse", string(data)}, " ")
}
