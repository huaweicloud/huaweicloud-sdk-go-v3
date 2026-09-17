package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelScrumAssociateResponse Response Object
type CancelScrumAssociateResponse struct {
	Result *CancelAssociateIssueResponseResult `json:"result,omitempty"`

	// **参数解释**： 接口整体响应状态。 **取值范围**： - success：取消关联工作项成功。 - error：取消关联工作项失败,详见错误码说明。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelScrumAssociateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelScrumAssociateResponse struct{}"
	}

	return strings.Join([]string{"CancelScrumAssociateResponse", string(data)}, " ")
}
