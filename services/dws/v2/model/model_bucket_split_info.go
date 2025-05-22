package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BucketSplitInfo **参数解释**： 桶分裂信息。 **取值范围**： 不涉及。
type BucketSplitInfo struct {

	// **参数解释**： 当前DN数。 **取值范围**： 不涉及。
	CurrentDnNum *int32 `json:"current_dn_num,omitempty"`

	// **参数解释**： 扩容后DN数。 **取值范围**： 不涉及。
	AfterScaleOutDnNum *int32 `json:"after_scale_out_dn_num,omitempty"`

	// **参数解释**： 当前bucket数。 **取值范围**： 不涉及。
	CurrentBucketNum *int32 `json:"current_bucket_num,omitempty"`

	// **参数解释**： 扩容后bucket数。 **取值范围**： 不涉及。
	AfterScaleOutBucketNum *int32 `json:"after_scale_out_bucket_num,omitempty"`

	// **参数解释**： 扩容是否涉及bucket分裂。 **取值范围**： 不涉及。
	IsBucketSplit *bool `json:"is_bucket_split,omitempty"`

	// **参数解释**： bucket DN倾斜率，用于衡量bucket在DN节点上不均衡程度。 **取值范围**： 不涉及。
	BucketTiltRate *string `json:"bucket_tilt_rate,omitempty"`

	// **参数解释**： 扩容后 bucket DN倾斜率，用于衡量扩容后bucket在DN节点上不均衡程度。 **取值范围**： 不涉及。
	AfterScaleOutBucketTiltRate *string `json:"after_scale_out_bucket_tilt_rate,omitempty"`
}

func (o BucketSplitInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BucketSplitInfo struct{}"
	}

	return strings.Join([]string{"BucketSplitInfo", string(data)}, " ")
}
