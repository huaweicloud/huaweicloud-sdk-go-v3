package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OpResourceExtraInfoResp struct {

	// **参数解释：** 需要排除备份的卷ID列表，列表元素为String类型的UUID。 **取值范围：** 不涉及
	ExcludeVolumes *[]string `json:"exclude_volumes,omitempty"`

	// **参数解释：** 指定需要备份的卷，如果有指定值，则每次备份都只备份指定的卷，如虚拟机绑定存储库之后新绑定的卷不会备份；如果为空默认为资源全部卷；仅虚拟机磁盘级备份特性中有效
	IncludeVolumes *[]ResourceExtraInfoIncludeVolumesResp `json:"include_volumes,omitempty"`
}

func (o OpResourceExtraInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpResourceExtraInfoResp struct{}"
	}

	return strings.Join([]string{"OpResourceExtraInfoResp", string(data)}, " ")
}
