package data

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/sparklee/abc-api/internal/conf"
	"github.com/sparklee/abc-api/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// init sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	lg := log.NewHelper(logger)

	client, err := makeEntClient(c, lg)
	if err != nil {
		return nil, nil, err
	}

	err = autoMigrate(client, lg)
	if err != nil {
		return nil, nil, err
	}

	return &Data{
			db: client,
		}, func() {
			lg.Info("closing the data resources")
			if err := client.Close(); err != nil {
				lg.Error(err)
			}
		}, nil
}

func makeEntClient(c *conf.Data, lg *log.Helper) (*ent.Client, error) {
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		lg.Errorf("failed opening connection to database: %v", err)
		return nil, err
	}
	sqlDrv := dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
		lg.WithContext(ctx).Info(i...)
	})
	client := ent.NewClient(ent.Driver(sqlDrv))
	return client, err
}

func autoMigrate(client *ent.Client, lg *log.Helper) error {
	if err := client.Schema.Create(context.Background()); err != nil {
		lg.Errorf("failed creating schema resources: %v", err)
		return err
	}
	return nil
}
