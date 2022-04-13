package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSecretStageRequestBody struct {
	// 凭据的版本号标识符。

	VersionId string `json:"version_id"`
}

func (o UpdateSecretStageRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecretStageRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSecretStageRequestBody", string(data)}, " ")
}
