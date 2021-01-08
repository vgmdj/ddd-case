package persistence

import (
	"context"
	"errors"

	"github.com/vgmdj/ddd-case/domain/manage/repository/facade"
	"github.com/vgmdj/ddd-case/infrastructure/repository/ent"
)

var (
	// ErrorNil if not register and use the global repo, will panic with ErrorNil
	ErrorNil = errors.New("must register first")
)

const (
	// ManagePersistenceRepo manage repo
	ManagePersistenceRepo = "manage"
)

// Config the config of persistence
type Config struct {
	Dialect string
	Source  string
}

// Repo the service of persistence
type Repo struct {
	manage facade.IManage
}

var (
	// GlobalRepo default global repo
	GlobalRepo *Repo
)

// NewPersistenceRepo return the factory of persistence factory
func NewPersistenceRepo(c *Config) *Repo {
	cli, err := ent.Open(c.Dialect, c.Source)
	if err != nil {
		panic(err.Error())
	}

	return &Repo{
		manage: NewManageRepository(cli),
	}
}

// InitGlobalRepo init the global repo
func InitGlobalRepo(c *Config) {
	GlobalRepo = NewPersistenceRepo(c)
}

// Ping check the database conn is ok or not
func (r *Repo) Ping(ctx context.Context) error {
	err := r.manage.Ping(ctx)
	if err != nil {
		return err
	}

	return nil
}

// ManagePersistenceRepo return the interface of manage
func (r *Repo) ManagePersistenceRepo() facade.IManage {
	return r.manage
}
