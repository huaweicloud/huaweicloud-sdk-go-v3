package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRepositoryByUuidResponse struct {
	Error *Error `json:"error,omitempty" xml:"error"`

	Result *RepoInfoV2 `json:"result,omitempty" xml:"result"`

	// 响应状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepositoryByUuidResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryByUuidResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryByUuidResponse", string(data)}, " ")
}
