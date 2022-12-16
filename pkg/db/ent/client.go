// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/detail"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/general"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/miningprofitdetail"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/miningprofitgeneral"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/miningprofitunsold"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/profit"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/withdraw"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Detail is the client for interacting with the Detail builders.
	Detail *DetailClient
	// General is the client for interacting with the General builders.
	General *GeneralClient
	// MiningProfitDetail is the client for interacting with the MiningProfitDetail builders.
	MiningProfitDetail *MiningProfitDetailClient
	// MiningProfitGeneral is the client for interacting with the MiningProfitGeneral builders.
	MiningProfitGeneral *MiningProfitGeneralClient
	// MiningProfitUnsold is the client for interacting with the MiningProfitUnsold builders.
	MiningProfitUnsold *MiningProfitUnsoldClient
	// Profit is the client for interacting with the Profit builders.
	Profit *ProfitClient
	// Withdraw is the client for interacting with the Withdraw builders.
	Withdraw *WithdrawClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Detail = NewDetailClient(c.config)
	c.General = NewGeneralClient(c.config)
	c.MiningProfitDetail = NewMiningProfitDetailClient(c.config)
	c.MiningProfitGeneral = NewMiningProfitGeneralClient(c.config)
	c.MiningProfitUnsold = NewMiningProfitUnsoldClient(c.config)
	c.Profit = NewProfitClient(c.config)
	c.Withdraw = NewWithdrawClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Detail:              NewDetailClient(cfg),
		General:             NewGeneralClient(cfg),
		MiningProfitDetail:  NewMiningProfitDetailClient(cfg),
		MiningProfitGeneral: NewMiningProfitGeneralClient(cfg),
		MiningProfitUnsold:  NewMiningProfitUnsoldClient(cfg),
		Profit:              NewProfitClient(cfg),
		Withdraw:            NewWithdrawClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		Detail:              NewDetailClient(cfg),
		General:             NewGeneralClient(cfg),
		MiningProfitDetail:  NewMiningProfitDetailClient(cfg),
		MiningProfitGeneral: NewMiningProfitGeneralClient(cfg),
		MiningProfitUnsold:  NewMiningProfitUnsoldClient(cfg),
		Profit:              NewProfitClient(cfg),
		Withdraw:            NewWithdrawClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Detail.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Detail.Use(hooks...)
	c.General.Use(hooks...)
	c.MiningProfitDetail.Use(hooks...)
	c.MiningProfitGeneral.Use(hooks...)
	c.MiningProfitUnsold.Use(hooks...)
	c.Profit.Use(hooks...)
	c.Withdraw.Use(hooks...)
}

// DetailClient is a client for the Detail schema.
type DetailClient struct {
	config
}

// NewDetailClient returns a client for the Detail from the given config.
func NewDetailClient(c config) *DetailClient {
	return &DetailClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `detail.Hooks(f(g(h())))`.
func (c *DetailClient) Use(hooks ...Hook) {
	c.hooks.Detail = append(c.hooks.Detail, hooks...)
}

// Create returns a builder for creating a Detail entity.
func (c *DetailClient) Create() *DetailCreate {
	mutation := newDetailMutation(c.config, OpCreate)
	return &DetailCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Detail entities.
func (c *DetailClient) CreateBulk(builders ...*DetailCreate) *DetailCreateBulk {
	return &DetailCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Detail.
func (c *DetailClient) Update() *DetailUpdate {
	mutation := newDetailMutation(c.config, OpUpdate)
	return &DetailUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DetailClient) UpdateOne(d *Detail) *DetailUpdateOne {
	mutation := newDetailMutation(c.config, OpUpdateOne, withDetail(d))
	return &DetailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DetailClient) UpdateOneID(id uuid.UUID) *DetailUpdateOne {
	mutation := newDetailMutation(c.config, OpUpdateOne, withDetailID(id))
	return &DetailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Detail.
func (c *DetailClient) Delete() *DetailDelete {
	mutation := newDetailMutation(c.config, OpDelete)
	return &DetailDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DetailClient) DeleteOne(d *Detail) *DetailDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *DetailClient) DeleteOneID(id uuid.UUID) *DetailDeleteOne {
	builder := c.Delete().Where(detail.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DetailDeleteOne{builder}
}

// Query returns a query builder for Detail.
func (c *DetailClient) Query() *DetailQuery {
	return &DetailQuery{
		config: c.config,
	}
}

// Get returns a Detail entity by its id.
func (c *DetailClient) Get(ctx context.Context, id uuid.UUID) (*Detail, error) {
	return c.Query().Where(detail.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DetailClient) GetX(ctx context.Context, id uuid.UUID) *Detail {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DetailClient) Hooks() []Hook {
	hooks := c.hooks.Detail
	return append(hooks[:len(hooks):len(hooks)], detail.Hooks[:]...)
}

// GeneralClient is a client for the General schema.
type GeneralClient struct {
	config
}

// NewGeneralClient returns a client for the General from the given config.
func NewGeneralClient(c config) *GeneralClient {
	return &GeneralClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `general.Hooks(f(g(h())))`.
func (c *GeneralClient) Use(hooks ...Hook) {
	c.hooks.General = append(c.hooks.General, hooks...)
}

// Create returns a builder for creating a General entity.
func (c *GeneralClient) Create() *GeneralCreate {
	mutation := newGeneralMutation(c.config, OpCreate)
	return &GeneralCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of General entities.
func (c *GeneralClient) CreateBulk(builders ...*GeneralCreate) *GeneralCreateBulk {
	return &GeneralCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for General.
func (c *GeneralClient) Update() *GeneralUpdate {
	mutation := newGeneralMutation(c.config, OpUpdate)
	return &GeneralUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GeneralClient) UpdateOne(ge *General) *GeneralUpdateOne {
	mutation := newGeneralMutation(c.config, OpUpdateOne, withGeneral(ge))
	return &GeneralUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GeneralClient) UpdateOneID(id uuid.UUID) *GeneralUpdateOne {
	mutation := newGeneralMutation(c.config, OpUpdateOne, withGeneralID(id))
	return &GeneralUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for General.
func (c *GeneralClient) Delete() *GeneralDelete {
	mutation := newGeneralMutation(c.config, OpDelete)
	return &GeneralDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GeneralClient) DeleteOne(ge *General) *GeneralDeleteOne {
	return c.DeleteOneID(ge.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *GeneralClient) DeleteOneID(id uuid.UUID) *GeneralDeleteOne {
	builder := c.Delete().Where(general.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GeneralDeleteOne{builder}
}

// Query returns a query builder for General.
func (c *GeneralClient) Query() *GeneralQuery {
	return &GeneralQuery{
		config: c.config,
	}
}

// Get returns a General entity by its id.
func (c *GeneralClient) Get(ctx context.Context, id uuid.UUID) (*General, error) {
	return c.Query().Where(general.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GeneralClient) GetX(ctx context.Context, id uuid.UUID) *General {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GeneralClient) Hooks() []Hook {
	hooks := c.hooks.General
	return append(hooks[:len(hooks):len(hooks)], general.Hooks[:]...)
}

// MiningProfitDetailClient is a client for the MiningProfitDetail schema.
type MiningProfitDetailClient struct {
	config
}

// NewMiningProfitDetailClient returns a client for the MiningProfitDetail from the given config.
func NewMiningProfitDetailClient(c config) *MiningProfitDetailClient {
	return &MiningProfitDetailClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `miningprofitdetail.Hooks(f(g(h())))`.
func (c *MiningProfitDetailClient) Use(hooks ...Hook) {
	c.hooks.MiningProfitDetail = append(c.hooks.MiningProfitDetail, hooks...)
}

// Create returns a builder for creating a MiningProfitDetail entity.
func (c *MiningProfitDetailClient) Create() *MiningProfitDetailCreate {
	mutation := newMiningProfitDetailMutation(c.config, OpCreate)
	return &MiningProfitDetailCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MiningProfitDetail entities.
func (c *MiningProfitDetailClient) CreateBulk(builders ...*MiningProfitDetailCreate) *MiningProfitDetailCreateBulk {
	return &MiningProfitDetailCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MiningProfitDetail.
func (c *MiningProfitDetailClient) Update() *MiningProfitDetailUpdate {
	mutation := newMiningProfitDetailMutation(c.config, OpUpdate)
	return &MiningProfitDetailUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MiningProfitDetailClient) UpdateOne(mpd *MiningProfitDetail) *MiningProfitDetailUpdateOne {
	mutation := newMiningProfitDetailMutation(c.config, OpUpdateOne, withMiningProfitDetail(mpd))
	return &MiningProfitDetailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MiningProfitDetailClient) UpdateOneID(id uuid.UUID) *MiningProfitDetailUpdateOne {
	mutation := newMiningProfitDetailMutation(c.config, OpUpdateOne, withMiningProfitDetailID(id))
	return &MiningProfitDetailUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MiningProfitDetail.
func (c *MiningProfitDetailClient) Delete() *MiningProfitDetailDelete {
	mutation := newMiningProfitDetailMutation(c.config, OpDelete)
	return &MiningProfitDetailDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MiningProfitDetailClient) DeleteOne(mpd *MiningProfitDetail) *MiningProfitDetailDeleteOne {
	return c.DeleteOneID(mpd.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MiningProfitDetailClient) DeleteOneID(id uuid.UUID) *MiningProfitDetailDeleteOne {
	builder := c.Delete().Where(miningprofitdetail.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MiningProfitDetailDeleteOne{builder}
}

// Query returns a query builder for MiningProfitDetail.
func (c *MiningProfitDetailClient) Query() *MiningProfitDetailQuery {
	return &MiningProfitDetailQuery{
		config: c.config,
	}
}

// Get returns a MiningProfitDetail entity by its id.
func (c *MiningProfitDetailClient) Get(ctx context.Context, id uuid.UUID) (*MiningProfitDetail, error) {
	return c.Query().Where(miningprofitdetail.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MiningProfitDetailClient) GetX(ctx context.Context, id uuid.UUID) *MiningProfitDetail {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MiningProfitDetailClient) Hooks() []Hook {
	hooks := c.hooks.MiningProfitDetail
	return append(hooks[:len(hooks):len(hooks)], miningprofitdetail.Hooks[:]...)
}

// MiningProfitGeneralClient is a client for the MiningProfitGeneral schema.
type MiningProfitGeneralClient struct {
	config
}

// NewMiningProfitGeneralClient returns a client for the MiningProfitGeneral from the given config.
func NewMiningProfitGeneralClient(c config) *MiningProfitGeneralClient {
	return &MiningProfitGeneralClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `miningprofitgeneral.Hooks(f(g(h())))`.
func (c *MiningProfitGeneralClient) Use(hooks ...Hook) {
	c.hooks.MiningProfitGeneral = append(c.hooks.MiningProfitGeneral, hooks...)
}

// Create returns a builder for creating a MiningProfitGeneral entity.
func (c *MiningProfitGeneralClient) Create() *MiningProfitGeneralCreate {
	mutation := newMiningProfitGeneralMutation(c.config, OpCreate)
	return &MiningProfitGeneralCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MiningProfitGeneral entities.
func (c *MiningProfitGeneralClient) CreateBulk(builders ...*MiningProfitGeneralCreate) *MiningProfitGeneralCreateBulk {
	return &MiningProfitGeneralCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MiningProfitGeneral.
func (c *MiningProfitGeneralClient) Update() *MiningProfitGeneralUpdate {
	mutation := newMiningProfitGeneralMutation(c.config, OpUpdate)
	return &MiningProfitGeneralUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MiningProfitGeneralClient) UpdateOne(mpg *MiningProfitGeneral) *MiningProfitGeneralUpdateOne {
	mutation := newMiningProfitGeneralMutation(c.config, OpUpdateOne, withMiningProfitGeneral(mpg))
	return &MiningProfitGeneralUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MiningProfitGeneralClient) UpdateOneID(id uuid.UUID) *MiningProfitGeneralUpdateOne {
	mutation := newMiningProfitGeneralMutation(c.config, OpUpdateOne, withMiningProfitGeneralID(id))
	return &MiningProfitGeneralUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MiningProfitGeneral.
func (c *MiningProfitGeneralClient) Delete() *MiningProfitGeneralDelete {
	mutation := newMiningProfitGeneralMutation(c.config, OpDelete)
	return &MiningProfitGeneralDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MiningProfitGeneralClient) DeleteOne(mpg *MiningProfitGeneral) *MiningProfitGeneralDeleteOne {
	return c.DeleteOneID(mpg.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MiningProfitGeneralClient) DeleteOneID(id uuid.UUID) *MiningProfitGeneralDeleteOne {
	builder := c.Delete().Where(miningprofitgeneral.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MiningProfitGeneralDeleteOne{builder}
}

// Query returns a query builder for MiningProfitGeneral.
func (c *MiningProfitGeneralClient) Query() *MiningProfitGeneralQuery {
	return &MiningProfitGeneralQuery{
		config: c.config,
	}
}

// Get returns a MiningProfitGeneral entity by its id.
func (c *MiningProfitGeneralClient) Get(ctx context.Context, id uuid.UUID) (*MiningProfitGeneral, error) {
	return c.Query().Where(miningprofitgeneral.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MiningProfitGeneralClient) GetX(ctx context.Context, id uuid.UUID) *MiningProfitGeneral {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MiningProfitGeneralClient) Hooks() []Hook {
	hooks := c.hooks.MiningProfitGeneral
	return append(hooks[:len(hooks):len(hooks)], miningprofitgeneral.Hooks[:]...)
}

// MiningProfitUnsoldClient is a client for the MiningProfitUnsold schema.
type MiningProfitUnsoldClient struct {
	config
}

// NewMiningProfitUnsoldClient returns a client for the MiningProfitUnsold from the given config.
func NewMiningProfitUnsoldClient(c config) *MiningProfitUnsoldClient {
	return &MiningProfitUnsoldClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `miningprofitunsold.Hooks(f(g(h())))`.
func (c *MiningProfitUnsoldClient) Use(hooks ...Hook) {
	c.hooks.MiningProfitUnsold = append(c.hooks.MiningProfitUnsold, hooks...)
}

// Create returns a builder for creating a MiningProfitUnsold entity.
func (c *MiningProfitUnsoldClient) Create() *MiningProfitUnsoldCreate {
	mutation := newMiningProfitUnsoldMutation(c.config, OpCreate)
	return &MiningProfitUnsoldCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of MiningProfitUnsold entities.
func (c *MiningProfitUnsoldClient) CreateBulk(builders ...*MiningProfitUnsoldCreate) *MiningProfitUnsoldCreateBulk {
	return &MiningProfitUnsoldCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for MiningProfitUnsold.
func (c *MiningProfitUnsoldClient) Update() *MiningProfitUnsoldUpdate {
	mutation := newMiningProfitUnsoldMutation(c.config, OpUpdate)
	return &MiningProfitUnsoldUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MiningProfitUnsoldClient) UpdateOne(mpu *MiningProfitUnsold) *MiningProfitUnsoldUpdateOne {
	mutation := newMiningProfitUnsoldMutation(c.config, OpUpdateOne, withMiningProfitUnsold(mpu))
	return &MiningProfitUnsoldUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MiningProfitUnsoldClient) UpdateOneID(id uuid.UUID) *MiningProfitUnsoldUpdateOne {
	mutation := newMiningProfitUnsoldMutation(c.config, OpUpdateOne, withMiningProfitUnsoldID(id))
	return &MiningProfitUnsoldUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for MiningProfitUnsold.
func (c *MiningProfitUnsoldClient) Delete() *MiningProfitUnsoldDelete {
	mutation := newMiningProfitUnsoldMutation(c.config, OpDelete)
	return &MiningProfitUnsoldDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MiningProfitUnsoldClient) DeleteOne(mpu *MiningProfitUnsold) *MiningProfitUnsoldDeleteOne {
	return c.DeleteOneID(mpu.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *MiningProfitUnsoldClient) DeleteOneID(id uuid.UUID) *MiningProfitUnsoldDeleteOne {
	builder := c.Delete().Where(miningprofitunsold.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MiningProfitUnsoldDeleteOne{builder}
}

// Query returns a query builder for MiningProfitUnsold.
func (c *MiningProfitUnsoldClient) Query() *MiningProfitUnsoldQuery {
	return &MiningProfitUnsoldQuery{
		config: c.config,
	}
}

// Get returns a MiningProfitUnsold entity by its id.
func (c *MiningProfitUnsoldClient) Get(ctx context.Context, id uuid.UUID) (*MiningProfitUnsold, error) {
	return c.Query().Where(miningprofitunsold.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MiningProfitUnsoldClient) GetX(ctx context.Context, id uuid.UUID) *MiningProfitUnsold {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *MiningProfitUnsoldClient) Hooks() []Hook {
	hooks := c.hooks.MiningProfitUnsold
	return append(hooks[:len(hooks):len(hooks)], miningprofitunsold.Hooks[:]...)
}

// ProfitClient is a client for the Profit schema.
type ProfitClient struct {
	config
}

// NewProfitClient returns a client for the Profit from the given config.
func NewProfitClient(c config) *ProfitClient {
	return &ProfitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `profit.Hooks(f(g(h())))`.
func (c *ProfitClient) Use(hooks ...Hook) {
	c.hooks.Profit = append(c.hooks.Profit, hooks...)
}

// Create returns a builder for creating a Profit entity.
func (c *ProfitClient) Create() *ProfitCreate {
	mutation := newProfitMutation(c.config, OpCreate)
	return &ProfitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Profit entities.
func (c *ProfitClient) CreateBulk(builders ...*ProfitCreate) *ProfitCreateBulk {
	return &ProfitCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Profit.
func (c *ProfitClient) Update() *ProfitUpdate {
	mutation := newProfitMutation(c.config, OpUpdate)
	return &ProfitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProfitClient) UpdateOne(pr *Profit) *ProfitUpdateOne {
	mutation := newProfitMutation(c.config, OpUpdateOne, withProfit(pr))
	return &ProfitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProfitClient) UpdateOneID(id uuid.UUID) *ProfitUpdateOne {
	mutation := newProfitMutation(c.config, OpUpdateOne, withProfitID(id))
	return &ProfitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Profit.
func (c *ProfitClient) Delete() *ProfitDelete {
	mutation := newProfitMutation(c.config, OpDelete)
	return &ProfitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProfitClient) DeleteOne(pr *Profit) *ProfitDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ProfitClient) DeleteOneID(id uuid.UUID) *ProfitDeleteOne {
	builder := c.Delete().Where(profit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProfitDeleteOne{builder}
}

// Query returns a query builder for Profit.
func (c *ProfitClient) Query() *ProfitQuery {
	return &ProfitQuery{
		config: c.config,
	}
}

// Get returns a Profit entity by its id.
func (c *ProfitClient) Get(ctx context.Context, id uuid.UUID) (*Profit, error) {
	return c.Query().Where(profit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProfitClient) GetX(ctx context.Context, id uuid.UUID) *Profit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ProfitClient) Hooks() []Hook {
	hooks := c.hooks.Profit
	return append(hooks[:len(hooks):len(hooks)], profit.Hooks[:]...)
}

// WithdrawClient is a client for the Withdraw schema.
type WithdrawClient struct {
	config
}

// NewWithdrawClient returns a client for the Withdraw from the given config.
func NewWithdrawClient(c config) *WithdrawClient {
	return &WithdrawClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `withdraw.Hooks(f(g(h())))`.
func (c *WithdrawClient) Use(hooks ...Hook) {
	c.hooks.Withdraw = append(c.hooks.Withdraw, hooks...)
}

// Create returns a builder for creating a Withdraw entity.
func (c *WithdrawClient) Create() *WithdrawCreate {
	mutation := newWithdrawMutation(c.config, OpCreate)
	return &WithdrawCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Withdraw entities.
func (c *WithdrawClient) CreateBulk(builders ...*WithdrawCreate) *WithdrawCreateBulk {
	return &WithdrawCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Withdraw.
func (c *WithdrawClient) Update() *WithdrawUpdate {
	mutation := newWithdrawMutation(c.config, OpUpdate)
	return &WithdrawUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WithdrawClient) UpdateOne(w *Withdraw) *WithdrawUpdateOne {
	mutation := newWithdrawMutation(c.config, OpUpdateOne, withWithdraw(w))
	return &WithdrawUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WithdrawClient) UpdateOneID(id uuid.UUID) *WithdrawUpdateOne {
	mutation := newWithdrawMutation(c.config, OpUpdateOne, withWithdrawID(id))
	return &WithdrawUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Withdraw.
func (c *WithdrawClient) Delete() *WithdrawDelete {
	mutation := newWithdrawMutation(c.config, OpDelete)
	return &WithdrawDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WithdrawClient) DeleteOne(w *Withdraw) *WithdrawDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *WithdrawClient) DeleteOneID(id uuid.UUID) *WithdrawDeleteOne {
	builder := c.Delete().Where(withdraw.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WithdrawDeleteOne{builder}
}

// Query returns a query builder for Withdraw.
func (c *WithdrawClient) Query() *WithdrawQuery {
	return &WithdrawQuery{
		config: c.config,
	}
}

// Get returns a Withdraw entity by its id.
func (c *WithdrawClient) Get(ctx context.Context, id uuid.UUID) (*Withdraw, error) {
	return c.Query().Where(withdraw.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WithdrawClient) GetX(ctx context.Context, id uuid.UUID) *Withdraw {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WithdrawClient) Hooks() []Hook {
	hooks := c.hooks.Withdraw
	return append(hooks[:len(hooks):len(hooks)], withdraw.Hooks[:]...)
}
