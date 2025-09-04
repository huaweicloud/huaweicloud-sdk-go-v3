package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetUserProfileReq 重置userprofile参数。
type ResetUserProfileReq struct {

	// id。
	CloudStorageAssignmentId string `json:"cloud_storage_assignment_id"`

	// 重置文件名称
	OriName *string `json:"ori_name,omitempty"`
}

func (o ResetUserProfileReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserProfileReq struct{}"
	}

	return strings.Join([]string{"ResetUserProfileReq", string(data)}, " ")
}
