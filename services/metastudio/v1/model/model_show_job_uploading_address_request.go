package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobUploadingAddressRequest Request Object
type ShowJobUploadingAddressRequest struct {

	// 任务id。
	JobId string `json:"job_id"`
}

func (o ShowJobUploadingAddressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobUploadingAddressRequest struct{}"
	}

	return strings.Join([]string{"ShowJobUploadingAddressRequest", string(data)}, " ")
}
