package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceExtraInfo struct {

	// 需要排除备份的卷id。仅在多系统盘备份特性中有效，排除不需要备份的磁盘。当虚拟机新绑定磁盘时，也能继续排除之前设置不用备份的卷。
	ExcludeVolumes *[]string `json:"exclude_volumes,omitempty" xml:"exclude_volumes"`

	// 指定需要备份的卷,若有指定值，则每次备份都只备份指定的卷，如虚拟机绑定存储库之后新绑定的卷不会备份；若为空默认为资源全部卷；仅虚拟机磁盘级备份特性中有效。
	IncludeVolumes *[]ResourceExtraInfoIncludeVolumes `json:"include_volumes,omitempty" xml:"include_volumes"`
}

func (o ResourceExtraInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceExtraInfo struct{}"
	}

	return strings.Join([]string{"ResourceExtraInfo", string(data)}, " ")
}
