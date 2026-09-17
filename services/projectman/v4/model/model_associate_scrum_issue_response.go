package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateScrumIssueResponse Response Object
type AssociateScrumIssueResponse struct {

	// **参数解释**： 本次新增的关联关系记录列表,每个元素对应一条关联关系。
	Result *[]AssociateIssueDetail `json:"result,omitempty"`

	// **参数解释**： 接口整体响应状态。 **取值范围**： - success：关联工作项成功。 - error：关联工作项失败,详见错误码说明。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AssociateScrumIssueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateScrumIssueResponse struct{}"
	}

	return strings.Join([]string{"AssociateScrumIssueResponse", string(data)}, " ")
}
