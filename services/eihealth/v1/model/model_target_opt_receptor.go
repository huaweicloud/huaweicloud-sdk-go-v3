package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TargetOptReceptor 靶点设置
type TargetOptReceptor struct {
	File *ReceptorDrugFile `json:"file"`

	// 是否平衡电荷
	BalancedCharge *bool `json:"balanced_charge,omitempty"`

	// 水模型, 支持选择spc, spce, tip3p, tip4p, tip5p
	WaterModel *string `json:"water_model,omitempty"`

	// 蛋白立场，支持选择amber03, amber94, amber96, amber99, amber99sb, amber99sb-ildn, amberGS, charmm27, oplsaa, gromos43a1, gromos43a2, gromos45a3, gromos53a5, gromos53a6, gromos54a7
	ForceField *string `json:"force_field,omitempty"`

	// 离子种类，支持选择NaCl、MgCl2、None，若设置了平衡电荷不支持选择None
	IonType *string `json:"ion_type,omitempty"`

	// 离子浓度，单位mol/L，若离子种类设置为None离子浓度不支持设置
	IonConcentration *float32 `json:"ion_concentration,omitempty"`
}

func (o TargetOptReceptor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetOptReceptor struct{}"
	}

	return strings.Join([]string{"TargetOptReceptor", string(data)}, " ")
}
