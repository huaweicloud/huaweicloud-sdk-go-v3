package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StatusSuccessResultWithUuidResult 结果
type StatusSuccessResultWithUuidResult struct {

	// uuid
	Uuid *string `json:"uuid,omitempty"`
}

func (o StatusSuccessResultWithUuidResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusSuccessResultWithUuidResult struct{}"
	}

	return strings.Join([]string{"StatusSuccessResultWithUuidResult", string(data)}, " ")
}
