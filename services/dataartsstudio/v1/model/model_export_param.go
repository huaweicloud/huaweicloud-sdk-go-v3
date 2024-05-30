package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExportParam struct {

	// 导出对象ID的列表，如：某几个逻辑模型的ID，填写String类型替代Long类型。
	Ids *[]string `json:"ids,omitempty"`

	// 是否是异步操作导出，true:异步，false:同步。
	Asyn *bool `json:"asyn,omitempty"`

	// 导出的业务类型：ER(关系建模)，Directory_CodeTable(码表目录)，Directory_Standard(标准目录)，DIM(维度建模)，codeTable(码表);dataStandard 数据标准;directory_id导出指定目录下的码表/数据标准;model_id，biz_catalog_id导出指定模型，目录下的业务表，import_bizcatalog导出流程架构，import_bizmetric导出业务指标。
	Type *string `json:"type,omitempty"`

	// 所属目录ID，填写String类型替代Long类型。
	DirectoryId *string `json:"directory_id,omitempty"`

	// 所属业务分层的ID，填写String类型替代Long类型。
	BizCatalogId *string `json:"biz_catalog_id,omitempty"`

	// 所属的业务分层的ID列表。
	BizCatalogIdList *[]string `json:"biz_catalog_id_list,omitempty"`

	// 所属关系建模的模型ID，导出关系模型需要此参数，填写String类型替代Long类型。
	ModelId *string `json:"model_id,omitempty"`
}

func (o ExportParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportParam struct{}"
	}

	return strings.Join([]string{"ExportParam", string(data)}, " ")
}
