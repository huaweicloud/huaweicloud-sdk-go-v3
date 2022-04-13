package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Severity struct {
	// 严重性等级取值范围：TIPS、LOW、MEDIUM、HIGH、FATAL。 TIPS：未发现任何问题。 LOW：无需针对问题执行任何操作。 MEDIUM：问题需要处理，但不紧急。 HIGH：问题必须优先处理。 FATAL：问题必须立即处理，以防止产生进一步的损害。

	Label string `json:"label"`
	// 严重性评分取值范围：0-100； 与严重性等级的对应关系： TIPS 0； LOW 1-39； MEDIUM 40-69； HIGH 70-89； FATAL 90-100。

	NormalizeScore *int32 `json:"normalize_score,omitempty"`
	// 严重性原始评分，指在数据源产品中的评分。

	OriginalScore *int32 `json:"original_score,omitempty"`
}

func (o Severity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Severity struct{}"
	}

	return strings.Join([]string{"Severity", string(data)}, " ")
}
