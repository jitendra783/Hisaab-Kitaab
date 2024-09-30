package db

import (
	"context"
	"database/sql"
	"hisaab-kitaab/pkg/config"
	e "hisaab-kitaab/pkg/errors"
	elog "hisaab-kitaab/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	// "gorm.io/gorm/schema"

	"time"
)

type dbLogger struct{}

func MysqlConnect() (*gorm.DB, *sql.DB, error) {
	c := config.GetConfig()

	// dsn := "host=" + c.GetString("db.mysql.host") +
	// 	" user=" + c.GetString("db.mysql.user") +
	// 	" password=" + c.GetString("db.mysql.password") +
	// 	" dbname=" + c.GetString("db.mysql.db") +
	// 	" port=" + c.GetString("db.mysql.port") +
	// 	" sslmode= " + c.GetString("db.mysql.sslmode")
	// // Connect to the database
	// mysql, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	elog.Log().Error("failed to connect mysql connection", zap.Error(err), zap.String("connStr", dsn))
	// 	return nil, mysql, e.ErrorInfo["MysqlDBConnError"]
	// }

	dsn1 := "host=" + c.GetString("db.postgres.host") +
		" user=" + c.GetString("db.postgres.user") +
		" password=" + c.GetString("db.postgres.password") +
		" dbname=" + c.GetString("db.postgres.db") +
		" port=" + c.GetString("db.postgres.port") +
		" sslmode= " + c.GetString("db.postgres.sslmode")

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn1), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: customLogger(),
	})
	if err != nil {
		elog.Log().Error("failed to connect postgreSQL connection", zap.Error(err), zap.String("connStr", dsn1))
		return nil, nil, e.ErrorInfo["postreSQLDBConnError"]
	}
	elog.Log().Info("postgre Database Connected")
	return db, nil, nil
}

func customLogger() logger.Interface {
	return dbLogger{}
}

func (d dbLogger) Error(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("error", data), zap.Any("description", others))
}

func (d dbLogger) Info(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("msg", data), zap.Any("description", others))
}

func (d dbLogger) Warn(ctx context.Context, data string, others ...interface{}) {
	elog.Log(ctx).Info("database", zap.String("msg", data), zap.Any("description", others))
}

func (d dbLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	query, others := fc()
	if err != nil {
		elog.Log(ctx).Info("database", zap.String("query", query), zap.Any("rows-affected", others), zap.Error(err))
	} else {
		elog.Log(ctx).Info("database", zap.String("query", query), zap.Any("rows-affected", others))
	}
}

func (d dbLogger) LogMode(l logger.LogLevel) logger.Interface {
	return d
}
