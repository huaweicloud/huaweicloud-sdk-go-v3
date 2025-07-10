package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeystorePermissionRequestBody 执行任务接口请求体
type UpdateKeystorePermissionRequestBody struct {

	// can_absent
	CanAbsent *bool `json:"can_absent,omitempty"`

	// delete
	Delete bool `json:"delete"`

	// id
	Id string `json:"id"`

	// keystore_id
	KeystoreId string `json:"keystore_id"`

	// modify
	Modify bool `json:"modify"`

	// usage
	Usage bool `json:"usage"`

	// user_name
	UserName string `json:"user_name"`
}

func (o UpdateKeystorePermissionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeystorePermissionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKeystorePermissionRequestBody", string(data)}, " ")
}
