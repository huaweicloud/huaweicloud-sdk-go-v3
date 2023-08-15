package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDeployKeyRequest Request Object
type AddDeployKeyRequest struct {

	// 仓库短id
	RepositoryId int32 `json:"repository_id"`

	Body *AddDeployKeyRequestBody `json:"body,omitempty"`
}

func (o AddDeployKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeployKeyRequest struct{}"
	}

	return strings.Join([]string{"AddDeployKeyRequest", string(data)}, " ")
}
