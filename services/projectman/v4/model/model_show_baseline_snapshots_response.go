package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaselineSnapshotsResponse Response Object
type ShowBaselineSnapshotsResponse struct {

	// **参数解释**： 返回状态。 **取值范围**： - success：响应成功 - error：响应失败
	Status *string `json:"status,omitempty"`

	// 消息
	Message *string `json:"message,omitempty"`

	// 返回结果
	Result         *[]FeatureSetOpenApiVo `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowBaselineSnapshotsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineSnapshotsResponse struct{}"
	}

	return strings.Join([]string{"ShowBaselineSnapshotsResponse", string(data)}, " ")
}
