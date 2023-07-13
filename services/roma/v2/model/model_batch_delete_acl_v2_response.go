package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAclV2Response Response Object
type BatchDeleteAclV2Response struct {

	// 删除成功的ACL策略数量
	SuccessCount *int32 `json:"success_count,omitempty"`

	// 删除失败的ACL策略及错误信息
	Failure        *[]AclBatchFailure `json:"failure,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o BatchDeleteAclV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAclV2Response struct{}"
	}

	return strings.Join([]string{"BatchDeleteAclV2Response", string(data)}, " ")
}
