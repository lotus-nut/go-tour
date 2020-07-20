package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DBTypeToStructType = map[string]string{
	"tinyint":   "int8",
	"smallint":  "int",
	"int":       "int32",
	"mediumint": "int64",
	"bigint":    "int64",
	"bool":      "bool",
	"varchar":   "string",
	"text":      "string",
	"date":      "string",
	"datetime":  "string",
}

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/information_schema?charset=%s&parseTime&loc=Local",
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := "SELECT " +
		"COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS " +
		"WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey,
			&column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}
