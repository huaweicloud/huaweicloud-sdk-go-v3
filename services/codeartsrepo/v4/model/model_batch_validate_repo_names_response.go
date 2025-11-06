package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchValidateRepoNamesResponse Response Object
type BatchValidateRepoNamesResponse struct {

	// 检查结果
	Body           *[]ValidateRepoNameResultDto `json:"body,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o BatchValidateRepoNamesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchValidateRepoNamesResponse struct{}"
	}

	return strings.Join([]string{"BatchValidateRepoNamesResponse", string(data)}, " ")
}
