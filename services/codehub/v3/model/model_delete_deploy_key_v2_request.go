package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDeployKeyV2Request struct {
	// 部署密钥id

	KeyId int32 `json:"key_id"`
	// 仓库主键id

	RepositoryId int32 `json:"repository_id"`
}

func (o DeleteDeployKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeployKeyV2Request struct{}"
	}

	return strings.Join([]string{"DeleteDeployKeyV2Request", string(data)}, " ")
}
