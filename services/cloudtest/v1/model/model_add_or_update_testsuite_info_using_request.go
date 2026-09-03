package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddOrUpdateTestsuiteInfoUsingRequest Request Object
type AddOrUpdateTestsuiteInfoUsingRequest struct {

	// 服务id
	ServiceId string `json:"service_id"`

	Body *TaskInfoV4VoReq `json:"body,omitempty"`
}

func (o AddOrUpdateTestsuiteInfoUsingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOrUpdateTestsuiteInfoUsingRequest struct{}"
	}

	return strings.Join([]string{"AddOrUpdateTestsuiteInfoUsingRequest", string(data)}, " ")
}
