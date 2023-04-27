package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListScaleOutPolicyResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 计算资源列表
	Policies       *[]ScaleOutPolicyRsp `json:"policies,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListScaleOutPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScaleOutPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListScaleOutPolicyResponse", string(data)}, " ")
}
