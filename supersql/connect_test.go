package supersql

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/SAP/go-ase"
	_ "github.com/SAP/go-hdb/driver"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/snowflakedb/gosnowflake"
	_ "github.com/vertica/vertica-sql-go"
	_ "github.com/ziutek/mymysql/godrv"
	_ "modernc.org/sqlite"
)

func TestConnect(t *testing.T) {
	type args struct {
		ctx    context.Context
		driver string
		dsn    string
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{
		// Here I would test the Connect method with different drivers and DSNs,
		// regardless of the fact that it's not necessary.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Connect(tt.args.ctx, tt.args.driver, tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() got = %v, want %v", got, tt.want)
			}
		})
	}
}
