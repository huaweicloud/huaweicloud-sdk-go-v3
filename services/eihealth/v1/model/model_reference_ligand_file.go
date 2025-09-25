package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReferenceLigandFile 模板配体文件，仅引擎为SIMILAR_DOCKING时提供。
type ReferenceLigandFile struct {
	Source *DrugFileSource `json:"source"`

	// **参数解释**： 文件URL。 **约束限制**： 当数据源source为外部网络数据时为https地址，为用户私有数据中心时为空间路径，为公共数据场景时为obs地址。 **取值范围**： 文件URL仅支持以.pdb、.sdf、.mol2、.smi、.csv结尾，长度为[1-2000]个字符。 **默认取值**： 不涉及
	Url *string `json:"url,omitempty"`

	// **参数解释**： 文件格式。 **约束限制**： 仅数据源source为RAW时提供。 **取值范围**： - PDB - SDF - MOL2 - SMI - CSV **默认取值**： 不涉及
	Format *string `json:"format,omitempty"`

	// **参数解释**： 文件原始数据。 **约束限制**： 仅数据源source为RAW时提供。 **取值范围**： 长度为[0-10000000]个字符。 **默认取值**： 不涉及
	Data *string `json:"data,omitempty"`
}

func (o ReferenceLigandFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReferenceLigandFile struct{}"
	}

	return strings.Join([]string{"ReferenceLigandFile", string(data)}, " ")
}
