package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SqlTypeRangeConfigResult 全量SQL开关记录
type SqlTypeRangeConfigResult struct {

	// **参数解释**: SQL类型的归类名称。 对常用SQL类型，简单归类为6个类别，每个类别对应一组固定的采集SQL语句类型列表，采用前缀进行匹配。 对于其他场景，可以使用自定义类别，允许按需自定义采集SQL的语句前缀。 **取值范围**: - all：不区分SQL类型，全部采集。 - ddl：只采集DDL语句类型。 - dml：只采集DML语句类型。 - dcl：只采集DCL语句类型。 - tcl：只采集TCL语句类型。 - dql：只采集DQL语句类型。 - custom：采集自定义语句类型。
	Category *string `json:"category,omitempty"`

	// **参数解释**: 对应SQL类别中，采集的SQL语句类型列表，采用前缀方式进行匹配。 **取值范围**: - all：不区分SQL类型，全部采集。对应取值为：[\".*\"]。 - ddl：只采集DDL语句类别，对于取值为：[\"create\", \"alter\", \"drop\", \"truncate\", \"reindex\", \"vacuum\", \"analyze\", \"declare\", \"move\", \"close\"]。 - dml：只采集DML语句类型，对于取值为：[\"insert\", \"update\", \"delete\", \"merge\", \"show\", \"explain\", \"prepare\", \"lock\", \"copy\", \"execute\", \"deallocate\"]。 - dcl：只采集DCL语句类型，对于取值为：[\"grant\", \"revoke\", \"reassign\", \"set\"]。 - tcl：只采集TCL语句类型，对于取值为：[\"begin\", \"commit\", \"rollback\", \"start\", \"savepoint\", \"checkpoint\", \"release savepoint\"]。 - dql：只采集DQL语句类型，对于取值为：[\"select\"]。 - custom：采集自定义语句类型。对应取值为：开启全量SQL时，用户填写的自定义SQL语句类型列表。
	Prefixes *string `json:"prefixes,omitempty"`

	// **参数解释**: 对应SQL类别是否为预置类别。 **取值范围**: - true：对应category取值all、ddl 、dml 、dcl 、tcl 、dql 。 - false：对应category取值custom。
	IsPreset *bool `json:"is_preset,omitempty"`
}

func (o SqlTypeRangeConfigResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlTypeRangeConfigResult struct{}"
	}

	return strings.Join([]string{"SqlTypeRangeConfigResult", string(data)}, " ")
}
