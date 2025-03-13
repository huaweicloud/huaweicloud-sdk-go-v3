package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecretDetail struct {

	// 凭据ID
	Id string `json:"id"`

	// 凭证名字。
	Name string `json:"name"`

	// 当前凭据是否有更新版本。
	IfUpdateAvailable bool `json:"if_update_available"`

	// 凭据在DEW的状态。
	SecretStatus string `json:"secret_status"`

	// 凭据在CAE使用状态。
	Status string `json:"status"`

	// 当前使用的凭证版本号。
	VersionId string `json:"version_id"`

	// 当前版本凭证在dew的创建时间。
	ModifiedTime *int32 `json:"modified_time,omitempty"`
}

func (o SecretDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecretDetail struct{}"
	}

	return strings.Join([]string{"SecretDetail", string(data)}, " ")
}
