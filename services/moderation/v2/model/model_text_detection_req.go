package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type TextDetectionReq struct {
	// 检测场景。  当前支持的场景有默认场景和用户自定义场景：  - 默认场景为：     * politics：涉政     * porn：涉黄     * ad：广告     * abuse：辱骂     * contraband：违禁品     * flood：灌水   - 用户自定义场景为：自定义黑名单词库。  > - 自定义词库的创建和使用请参见[配置自定义词库](https://support.huaweicloud.com/api-moderation/moderation_03_0027.html)。 > - flood场景不支持使用自定义白名单词库。

	Categories *[]string `json:"categories,omitempty"`
	// 待检测的文本列表，目前暂时每次只支持传一个item。

	Items []TextDetectionItemsReq `json:"items"`
}

func (o TextDetectionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextDetectionReq struct{}"
	}

	return strings.Join([]string{"TextDetectionReq", string(data)}, " ")
}
