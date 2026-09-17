package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkitemConfigsResponse Response Object
type ListWorkitemConfigsResponse struct {

	// **参数解释：** 已关闭工作项只读模式。 **取值范围：** true：无法进行编辑或修改。 false：可以进行编辑或修改。
	ClosedWorkitemReadonlyMode *bool `json:"closed_workitem_readonly_mode,omitempty"`
	HttpStatusCode             int   `json:"-"`
}

func (o ListWorkitemConfigsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkitemConfigsResponse struct{}"
	}

	return strings.Join([]string{"ListWorkitemConfigsResponse", string(data)}, " ")
}
