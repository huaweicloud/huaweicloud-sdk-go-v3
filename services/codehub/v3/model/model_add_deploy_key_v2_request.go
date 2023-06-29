package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddDeployKeyV2Request Request Object
type AddDeployKeyV2Request struct {

	// 仓库主键id
	RepositoryId int32 `json:"repository_id"`

	Body *AddDeployKeyRequestBody `json:"body,omitempty"`
}

func (o AddDeployKeyV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDeployKeyV2Request struct{}"
	}

	return strings.Join([]string{"AddDeployKeyV2Request", string(data)}, " ")
}
