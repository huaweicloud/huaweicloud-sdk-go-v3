package model

import (
	"encoding/json"

	"strings"
)

// 数据库信息。
type PostgresqlDatabaseForCreation struct {
	// 数据库名称。 数据库名称长度可在1～63个字符之间，由字母、数字、或下划线组成，不能包含其他特殊字符，不能以“pg”和数字开头，且不能和RDS for PostgreSQL模板库重名。 RDS for PostgreSQL模板库包括postgres， template0 ，template1。

	Name string `json:"name"`
	// 数据库字符集。默认UTF8。

	CharacterSet string `json:"character_set"`
	// 数据库所属用户，缺省时默认是root，不能和系统用户重名，且必须是已存在的用户。  系统用户包括“rdsAdmin”,“ rdsMetric”, “rdsBackup”, “rdsRepl”,“ rdsProxy”, “rdsDdm”。

	Owner string `json:"owner"`
	// 数据库模板名称，可选值template0，template1。默认template1。

	Template *string `json:"template,omitempty"`
	// 数据库排序集。默认en_US.UTF-8。注意，不同的排序规则下，相同的比较其结果可能是不同的。例如，在en_US.utf8下， select 'a'>'A';为false，但在'C'下，select 'a'>'A';为true; 如果从oracle迁移到PostgreSQL，使用'C'才能得到一致的预期。支持的排序规则，可以查询系统表 pg_collation。

	LcCollate *string `json:"lc_collate,omitempty"`
	// 数据库分类集。默认en_US.UTF-8。

	LcCtype *string `json:"lc_ctype,omitempty"`
}

func (o PostgresqlDatabaseForCreation) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PostgresqlDatabaseForCreation struct{}"
	}

	return strings.Join([]string{"PostgresqlDatabaseForCreation", string(data)}, " ")
}