// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/withdraw"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// WithdrawUpdate is the builder for updating Withdraw entities.
type WithdrawUpdate struct {
	config
	hooks    []Hook
	mutation *WithdrawMutation
}

// Where appends a list predicates to the WithdrawUpdate builder.
func (wu *WithdrawUpdate) Where(ps ...predicate.Withdraw) *WithdrawUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetCreatedAt sets the "created_at" field.
func (wu *WithdrawUpdate) SetCreatedAt(u uint32) *WithdrawUpdate {
	wu.mutation.ResetCreatedAt()
	wu.mutation.SetCreatedAt(u)
	return wu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableCreatedAt(u *uint32) *WithdrawUpdate {
	if u != nil {
		wu.SetCreatedAt(*u)
	}
	return wu
}

// AddCreatedAt adds u to the "created_at" field.
func (wu *WithdrawUpdate) AddCreatedAt(u int32) *WithdrawUpdate {
	wu.mutation.AddCreatedAt(u)
	return wu
}

// SetUpdatedAt sets the "updated_at" field.
func (wu *WithdrawUpdate) SetUpdatedAt(u uint32) *WithdrawUpdate {
	wu.mutation.ResetUpdatedAt()
	wu.mutation.SetUpdatedAt(u)
	return wu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (wu *WithdrawUpdate) AddUpdatedAt(u int32) *WithdrawUpdate {
	wu.mutation.AddUpdatedAt(u)
	return wu
}

// SetDeletedAt sets the "deleted_at" field.
func (wu *WithdrawUpdate) SetDeletedAt(u uint32) *WithdrawUpdate {
	wu.mutation.ResetDeletedAt()
	wu.mutation.SetDeletedAt(u)
	return wu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableDeletedAt(u *uint32) *WithdrawUpdate {
	if u != nil {
		wu.SetDeletedAt(*u)
	}
	return wu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (wu *WithdrawUpdate) AddDeletedAt(u int32) *WithdrawUpdate {
	wu.mutation.AddDeletedAt(u)
	return wu
}

// SetAppID sets the "app_id" field.
func (wu *WithdrawUpdate) SetAppID(u uuid.UUID) *WithdrawUpdate {
	wu.mutation.SetAppID(u)
	return wu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableAppID(u *uuid.UUID) *WithdrawUpdate {
	if u != nil {
		wu.SetAppID(*u)
	}
	return wu
}

// ClearAppID clears the value of the "app_id" field.
func (wu *WithdrawUpdate) ClearAppID() *WithdrawUpdate {
	wu.mutation.ClearAppID()
	return wu
}

// SetUserID sets the "user_id" field.
func (wu *WithdrawUpdate) SetUserID(u uuid.UUID) *WithdrawUpdate {
	wu.mutation.SetUserID(u)
	return wu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableUserID(u *uuid.UUID) *WithdrawUpdate {
	if u != nil {
		wu.SetUserID(*u)
	}
	return wu
}

// ClearUserID clears the value of the "user_id" field.
func (wu *WithdrawUpdate) ClearUserID() *WithdrawUpdate {
	wu.mutation.ClearUserID()
	return wu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (wu *WithdrawUpdate) SetCoinTypeID(u uuid.UUID) *WithdrawUpdate {
	wu.mutation.SetCoinTypeID(u)
	return wu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableCoinTypeID(u *uuid.UUID) *WithdrawUpdate {
	if u != nil {
		wu.SetCoinTypeID(*u)
	}
	return wu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (wu *WithdrawUpdate) ClearCoinTypeID() *WithdrawUpdate {
	wu.mutation.ClearCoinTypeID()
	return wu
}

// SetAccountID sets the "account_id" field.
func (wu *WithdrawUpdate) SetAccountID(u uuid.UUID) *WithdrawUpdate {
	wu.mutation.SetAccountID(u)
	return wu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableAccountID(u *uuid.UUID) *WithdrawUpdate {
	if u != nil {
		wu.SetAccountID(*u)
	}
	return wu
}

// ClearAccountID clears the value of the "account_id" field.
func (wu *WithdrawUpdate) ClearAccountID() *WithdrawUpdate {
	wu.mutation.ClearAccountID()
	return wu
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (wu *WithdrawUpdate) SetPlatformTransactionID(u uuid.UUID) *WithdrawUpdate {
	wu.mutation.SetPlatformTransactionID(u)
	return wu
}

// SetNillablePlatformTransactionID sets the "platform_transaction_id" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillablePlatformTransactionID(u *uuid.UUID) *WithdrawUpdate {
	if u != nil {
		wu.SetPlatformTransactionID(*u)
	}
	return wu
}

// ClearPlatformTransactionID clears the value of the "platform_transaction_id" field.
func (wu *WithdrawUpdate) ClearPlatformTransactionID() *WithdrawUpdate {
	wu.mutation.ClearPlatformTransactionID()
	return wu
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (wu *WithdrawUpdate) SetChainTransactionID(s string) *WithdrawUpdate {
	wu.mutation.SetChainTransactionID(s)
	return wu
}

// SetNillableChainTransactionID sets the "chain_transaction_id" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableChainTransactionID(s *string) *WithdrawUpdate {
	if s != nil {
		wu.SetChainTransactionID(*s)
	}
	return wu
}

// ClearChainTransactionID clears the value of the "chain_transaction_id" field.
func (wu *WithdrawUpdate) ClearChainTransactionID() *WithdrawUpdate {
	wu.mutation.ClearChainTransactionID()
	return wu
}

// SetState sets the "state" field.
func (wu *WithdrawUpdate) SetState(s string) *WithdrawUpdate {
	wu.mutation.SetState(s)
	return wu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableState(s *string) *WithdrawUpdate {
	if s != nil {
		wu.SetState(*s)
	}
	return wu
}

// ClearState clears the value of the "state" field.
func (wu *WithdrawUpdate) ClearState() *WithdrawUpdate {
	wu.mutation.ClearState()
	return wu
}

// SetAmount sets the "amount" field.
func (wu *WithdrawUpdate) SetAmount(d decimal.Decimal) *WithdrawUpdate {
	wu.mutation.ResetAmount()
	wu.mutation.SetAmount(d)
	return wu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (wu *WithdrawUpdate) SetNillableAmount(d *decimal.Decimal) *WithdrawUpdate {
	if d != nil {
		wu.SetAmount(*d)
	}
	return wu
}

// AddAmount adds d to the "amount" field.
func (wu *WithdrawUpdate) AddAmount(d decimal.Decimal) *WithdrawUpdate {
	wu.mutation.AddAmount(d)
	return wu
}

// ClearAmount clears the value of the "amount" field.
func (wu *WithdrawUpdate) ClearAmount() *WithdrawUpdate {
	wu.mutation.ClearAmount()
	return wu
}

// Mutation returns the WithdrawMutation object of the builder.
func (wu *WithdrawUpdate) Mutation() *WithdrawMutation {
	return wu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WithdrawUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := wu.defaults(); err != nil {
		return 0, err
	}
	if len(wu.hooks) == 0 {
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithdrawMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WithdrawUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WithdrawUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WithdrawUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WithdrawUpdate) defaults() error {
	if _, ok := wu.mutation.UpdatedAt(); !ok {
		if withdraw.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized withdraw.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := withdraw.UpdateDefaultUpdatedAt()
		wu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (wu *WithdrawUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withdraw.Table,
			Columns: withdraw.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: withdraw.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldCreatedAt,
		})
	}
	if value, ok := wu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldCreatedAt,
		})
	}
	if value, ok := wu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldUpdatedAt,
		})
	}
	if value, ok := wu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldUpdatedAt,
		})
	}
	if value, ok := wu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldDeletedAt,
		})
	}
	if value, ok := wu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldDeletedAt,
		})
	}
	if value, ok := wu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: withdraw.FieldAppID,
		})
	}
	if wu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: withdraw.FieldAppID,
		})
	}
	if value, ok := wu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: withdraw.FieldUserID,
		})
	}
	if wu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: withdraw.FieldUserID,
		})
	}
	if value, ok := wu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: withdraw.FieldCoinTypeID,
		})
	}
	if wu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: withdraw.FieldCoinTypeID,
		})
	}
	if value, ok := wu.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: withdraw.FieldAccountID,
		})
	}
	if wu.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: withdraw.FieldAccountID,
		})
	}
	if value, ok := wu.mutation.PlatformTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: withdraw.FieldPlatformTransactionID,
		})
	}
	if wu.mutation.PlatformTransactionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: withdraw.FieldPlatformTransactionID,
		})
	}
	if value, ok := wu.mutation.ChainTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: withdraw.FieldChainTransactionID,
		})
	}
	if wu.mutation.ChainTransactionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: withdraw.FieldChainTransactionID,
		})
	}
	if value, ok := wu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: withdraw.FieldState,
		})
	}
	if wu.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: withdraw.FieldState,
		})
	}
	if value, ok := wu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: withdraw.FieldAmount,
		})
	}
	if value, ok := wu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: withdraw.FieldAmount,
		})
	}
	if wu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: withdraw.FieldAmount,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withdraw.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WithdrawUpdateOne is the builder for updating a single Withdraw entity.
type WithdrawUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WithdrawMutation
}

// SetCreatedAt sets the "created_at" field.
func (wuo *WithdrawUpdateOne) SetCreatedAt(u uint32) *WithdrawUpdateOne {
	wuo.mutation.ResetCreatedAt()
	wuo.mutation.SetCreatedAt(u)
	return wuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableCreatedAt(u *uint32) *WithdrawUpdateOne {
	if u != nil {
		wuo.SetCreatedAt(*u)
	}
	return wuo
}

// AddCreatedAt adds u to the "created_at" field.
func (wuo *WithdrawUpdateOne) AddCreatedAt(u int32) *WithdrawUpdateOne {
	wuo.mutation.AddCreatedAt(u)
	return wuo
}

// SetUpdatedAt sets the "updated_at" field.
func (wuo *WithdrawUpdateOne) SetUpdatedAt(u uint32) *WithdrawUpdateOne {
	wuo.mutation.ResetUpdatedAt()
	wuo.mutation.SetUpdatedAt(u)
	return wuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (wuo *WithdrawUpdateOne) AddUpdatedAt(u int32) *WithdrawUpdateOne {
	wuo.mutation.AddUpdatedAt(u)
	return wuo
}

// SetDeletedAt sets the "deleted_at" field.
func (wuo *WithdrawUpdateOne) SetDeletedAt(u uint32) *WithdrawUpdateOne {
	wuo.mutation.ResetDeletedAt()
	wuo.mutation.SetDeletedAt(u)
	return wuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableDeletedAt(u *uint32) *WithdrawUpdateOne {
	if u != nil {
		wuo.SetDeletedAt(*u)
	}
	return wuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (wuo *WithdrawUpdateOne) AddDeletedAt(u int32) *WithdrawUpdateOne {
	wuo.mutation.AddDeletedAt(u)
	return wuo
}

// SetAppID sets the "app_id" field.
func (wuo *WithdrawUpdateOne) SetAppID(u uuid.UUID) *WithdrawUpdateOne {
	wuo.mutation.SetAppID(u)
	return wuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableAppID(u *uuid.UUID) *WithdrawUpdateOne {
	if u != nil {
		wuo.SetAppID(*u)
	}
	return wuo
}

// ClearAppID clears the value of the "app_id" field.
func (wuo *WithdrawUpdateOne) ClearAppID() *WithdrawUpdateOne {
	wuo.mutation.ClearAppID()
	return wuo
}

// SetUserID sets the "user_id" field.
func (wuo *WithdrawUpdateOne) SetUserID(u uuid.UUID) *WithdrawUpdateOne {
	wuo.mutation.SetUserID(u)
	return wuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableUserID(u *uuid.UUID) *WithdrawUpdateOne {
	if u != nil {
		wuo.SetUserID(*u)
	}
	return wuo
}

// ClearUserID clears the value of the "user_id" field.
func (wuo *WithdrawUpdateOne) ClearUserID() *WithdrawUpdateOne {
	wuo.mutation.ClearUserID()
	return wuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (wuo *WithdrawUpdateOne) SetCoinTypeID(u uuid.UUID) *WithdrawUpdateOne {
	wuo.mutation.SetCoinTypeID(u)
	return wuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *WithdrawUpdateOne {
	if u != nil {
		wuo.SetCoinTypeID(*u)
	}
	return wuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (wuo *WithdrawUpdateOne) ClearCoinTypeID() *WithdrawUpdateOne {
	wuo.mutation.ClearCoinTypeID()
	return wuo
}

// SetAccountID sets the "account_id" field.
func (wuo *WithdrawUpdateOne) SetAccountID(u uuid.UUID) *WithdrawUpdateOne {
	wuo.mutation.SetAccountID(u)
	return wuo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableAccountID(u *uuid.UUID) *WithdrawUpdateOne {
	if u != nil {
		wuo.SetAccountID(*u)
	}
	return wuo
}

// ClearAccountID clears the value of the "account_id" field.
func (wuo *WithdrawUpdateOne) ClearAccountID() *WithdrawUpdateOne {
	wuo.mutation.ClearAccountID()
	return wuo
}

// SetPlatformTransactionID sets the "platform_transaction_id" field.
func (wuo *WithdrawUpdateOne) SetPlatformTransactionID(u uuid.UUID) *WithdrawUpdateOne {
	wuo.mutation.SetPlatformTransactionID(u)
	return wuo
}

// SetNillablePlatformTransactionID sets the "platform_transaction_id" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillablePlatformTransactionID(u *uuid.UUID) *WithdrawUpdateOne {
	if u != nil {
		wuo.SetPlatformTransactionID(*u)
	}
	return wuo
}

// ClearPlatformTransactionID clears the value of the "platform_transaction_id" field.
func (wuo *WithdrawUpdateOne) ClearPlatformTransactionID() *WithdrawUpdateOne {
	wuo.mutation.ClearPlatformTransactionID()
	return wuo
}

// SetChainTransactionID sets the "chain_transaction_id" field.
func (wuo *WithdrawUpdateOne) SetChainTransactionID(s string) *WithdrawUpdateOne {
	wuo.mutation.SetChainTransactionID(s)
	return wuo
}

// SetNillableChainTransactionID sets the "chain_transaction_id" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableChainTransactionID(s *string) *WithdrawUpdateOne {
	if s != nil {
		wuo.SetChainTransactionID(*s)
	}
	return wuo
}

// ClearChainTransactionID clears the value of the "chain_transaction_id" field.
func (wuo *WithdrawUpdateOne) ClearChainTransactionID() *WithdrawUpdateOne {
	wuo.mutation.ClearChainTransactionID()
	return wuo
}

// SetState sets the "state" field.
func (wuo *WithdrawUpdateOne) SetState(s string) *WithdrawUpdateOne {
	wuo.mutation.SetState(s)
	return wuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableState(s *string) *WithdrawUpdateOne {
	if s != nil {
		wuo.SetState(*s)
	}
	return wuo
}

// ClearState clears the value of the "state" field.
func (wuo *WithdrawUpdateOne) ClearState() *WithdrawUpdateOne {
	wuo.mutation.ClearState()
	return wuo
}

// SetAmount sets the "amount" field.
func (wuo *WithdrawUpdateOne) SetAmount(d decimal.Decimal) *WithdrawUpdateOne {
	wuo.mutation.ResetAmount()
	wuo.mutation.SetAmount(d)
	return wuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (wuo *WithdrawUpdateOne) SetNillableAmount(d *decimal.Decimal) *WithdrawUpdateOne {
	if d != nil {
		wuo.SetAmount(*d)
	}
	return wuo
}

// AddAmount adds d to the "amount" field.
func (wuo *WithdrawUpdateOne) AddAmount(d decimal.Decimal) *WithdrawUpdateOne {
	wuo.mutation.AddAmount(d)
	return wuo
}

// ClearAmount clears the value of the "amount" field.
func (wuo *WithdrawUpdateOne) ClearAmount() *WithdrawUpdateOne {
	wuo.mutation.ClearAmount()
	return wuo
}

// Mutation returns the WithdrawMutation object of the builder.
func (wuo *WithdrawUpdateOne) Mutation() *WithdrawMutation {
	return wuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WithdrawUpdateOne) Select(field string, fields ...string) *WithdrawUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Withdraw entity.
func (wuo *WithdrawUpdateOne) Save(ctx context.Context) (*Withdraw, error) {
	var (
		err  error
		node *Withdraw
	)
	if err := wuo.defaults(); err != nil {
		return nil, err
	}
	if len(wuo.hooks) == 0 {
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WithdrawMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WithdrawUpdateOne) SaveX(ctx context.Context) *Withdraw {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WithdrawUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WithdrawUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WithdrawUpdateOne) defaults() error {
	if _, ok := wuo.mutation.UpdatedAt(); !ok {
		if withdraw.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized withdraw.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := withdraw.UpdateDefaultUpdatedAt()
		wuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (wuo *WithdrawUpdateOne) sqlSave(ctx context.Context) (_node *Withdraw, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   withdraw.Table,
			Columns: withdraw.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: withdraw.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Withdraw.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, withdraw.FieldID)
		for _, f := range fields {
			if !withdraw.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != withdraw.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldCreatedAt,
		})
	}
	if value, ok := wuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldCreatedAt,
		})
	}
	if value, ok := wuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldUpdatedAt,
		})
	}
	if value, ok := wuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldUpdatedAt,
		})
	}
	if value, ok := wuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldDeletedAt,
		})
	}
	if value, ok := wuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: withdraw.FieldDeletedAt,
		})
	}
	if value, ok := wuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: withdraw.FieldAppID,
		})
	}
	if wuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: withdraw.FieldAppID,
		})
	}
	if value, ok := wuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: withdraw.FieldUserID,
		})
	}
	if wuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: withdraw.FieldUserID,
		})
	}
	if value, ok := wuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: withdraw.FieldCoinTypeID,
		})
	}
	if wuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: withdraw.FieldCoinTypeID,
		})
	}
	if value, ok := wuo.mutation.AccountID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: withdraw.FieldAccountID,
		})
	}
	if wuo.mutation.AccountIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: withdraw.FieldAccountID,
		})
	}
	if value, ok := wuo.mutation.PlatformTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: withdraw.FieldPlatformTransactionID,
		})
	}
	if wuo.mutation.PlatformTransactionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: withdraw.FieldPlatformTransactionID,
		})
	}
	if value, ok := wuo.mutation.ChainTransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: withdraw.FieldChainTransactionID,
		})
	}
	if wuo.mutation.ChainTransactionIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: withdraw.FieldChainTransactionID,
		})
	}
	if value, ok := wuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: withdraw.FieldState,
		})
	}
	if wuo.mutation.StateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: withdraw.FieldState,
		})
	}
	if value, ok := wuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: withdraw.FieldAmount,
		})
	}
	if value, ok := wuo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: withdraw.FieldAmount,
		})
	}
	if wuo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Column: withdraw.FieldAmount,
		})
	}
	_node = &Withdraw{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{withdraw.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
