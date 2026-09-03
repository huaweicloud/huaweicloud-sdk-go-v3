package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiskSpaceDiagnosisResult 诊断结果
type DiskSpaceDiagnosisResult struct {

	// **参数解释**：  诊断项编码。  **约束限制**：  不涉及。  **取值范围**：  - 1001（慢查询using filesort产生临时文件） - 1002（慢查询using temporary产生临时文件） - 1003（大事务产生binlog临时文件） - 1004（未使用ROW_FORMAT创建临时表） - 1005（使用ROW_FORMAT创建临时表） - 1006（Online DDL创建临时文件） - 1007（DDL产生临时日志） - 2001（长事务产生undo文件） - 2002（慢日志） - 2003（审计日志） - 2004（binlog） - 2005（relaylog） - 3001（数据文件） - 4001（执行时间长） - 4002（临时表类） - 4003（排序类） - 4004（DDL类）  **默认取值**：  不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**：  诊断详情。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Detail *string `json:"detail,omitempty"`

	// **参数解释**：  用户查询时间内的磁盘容量是否受该诊断项影响，1代表是，0代表否。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Affect *int32 `json:"affect,omitempty"`
}

func (o DiskSpaceDiagnosisResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskSpaceDiagnosisResult struct{}"
	}

	return strings.Join([]string{"DiskSpaceDiagnosisResult", string(data)}, " ")
}
