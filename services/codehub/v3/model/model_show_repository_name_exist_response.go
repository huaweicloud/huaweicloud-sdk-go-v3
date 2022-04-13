package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowRepositoryNameExistResponse struct {
	Error *Error `json:"error,omitempty"`
	// 响应结果

	Result *bool `json:"result,omitempty"`
	// 响应状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowRepositoryNameExistResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryNameExistResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryNameExistResponse", string(data)}, " ")
}
