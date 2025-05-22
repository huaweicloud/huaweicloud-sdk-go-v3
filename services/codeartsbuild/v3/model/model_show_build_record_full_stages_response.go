package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildRecordFullStagesResponse Response Object
type ShowBuildRecordFullStagesResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *FullStagesResult `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowBuildRecordFullStagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildRecordFullStagesResponse struct{}"
	}

	return strings.Join([]string{"ShowBuildRecordFullStagesResponse", string(data)}, " ")
}
