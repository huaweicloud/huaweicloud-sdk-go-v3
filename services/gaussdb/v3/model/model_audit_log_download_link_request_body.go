package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AuditLogDownloadLinkRequestBody struct {

	// **参数解释**：  审计日志ID列表。  **约束限制**：  限制50条以内。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Ids []string `json:"ids"`
}

func (o AuditLogDownloadLinkRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditLogDownloadLinkRequestBody struct{}"
	}

	return strings.Join([]string{"AuditLogDownloadLinkRequestBody", string(data)}, " ")
}
