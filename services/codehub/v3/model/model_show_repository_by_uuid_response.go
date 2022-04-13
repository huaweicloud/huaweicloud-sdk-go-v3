package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRepositoryByUuidResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *RepoInfoV2 `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepositoryByUuidResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryByUuidResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryByUuidResponse", string(data)}, " ")
}
