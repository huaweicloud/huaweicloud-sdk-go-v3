package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateOrUpdateSecretDetail struct {

	// 凭证名字。
	Name string `json:"name"`

	// 当前使用的凭证版本号。
	VersionId string `json:"version_id"`
}

func (o CreateOrUpdateSecretDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateSecretDetail struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateSecretDetail", string(data)}, " ")
}
