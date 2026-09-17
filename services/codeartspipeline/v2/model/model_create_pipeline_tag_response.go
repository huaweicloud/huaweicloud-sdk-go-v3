package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePipelineTagResponse Response Object
type CreatePipelineTagResponse struct {

	// **参数解释**： 操作是否成功。 **取值范围**： - true：操作成功。 - false：操作失败。
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CreatePipelineTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePipelineTagResponse struct{}"
	}

	return strings.Join([]string{"CreatePipelineTagResponse", string(data)}, " ")
}
