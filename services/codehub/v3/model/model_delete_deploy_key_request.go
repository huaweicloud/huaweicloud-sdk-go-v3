package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteDeployKeyRequest struct {

	// 部署密钥id
	KeyId int32 `json:"key_id"`

	// 仓库短id
	RepositoryId int32 `json:"repository_id"`
}

func (o DeleteDeployKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDeployKeyRequest struct{}"
	}

	return strings.Join([]string{"DeleteDeployKeyRequest", string(data)}, " ")
}
