// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/detail"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/general"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/profit"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/schema"
	"github.com/NpoolPlatform/ledger-manager/pkg/db/ent/withdraw"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	detailMixin := schema.Detail{}.Mixin()
	detail.Policy = privacy.NewPolicies(detailMixin[0], schema.Detail{})
	detail.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := detail.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	detailMixinFields0 := detailMixin[0].Fields()
	_ = detailMixinFields0
	detailFields := schema.Detail{}.Fields()
	_ = detailFields
	// detailDescCreatedAt is the schema descriptor for created_at field.
	detailDescCreatedAt := detailMixinFields0[0].Descriptor()
	// detail.DefaultCreatedAt holds the default value on creation for the created_at field.
	detail.DefaultCreatedAt = detailDescCreatedAt.Default.(func() uint32)
	// detailDescUpdatedAt is the schema descriptor for updated_at field.
	detailDescUpdatedAt := detailMixinFields0[1].Descriptor()
	// detail.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	detail.DefaultUpdatedAt = detailDescUpdatedAt.Default.(func() uint32)
	// detail.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	detail.UpdateDefaultUpdatedAt = detailDescUpdatedAt.UpdateDefault.(func() uint32)
	// detailDescDeletedAt is the schema descriptor for deleted_at field.
	detailDescDeletedAt := detailMixinFields0[2].Descriptor()
	// detail.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	detail.DefaultDeletedAt = detailDescDeletedAt.Default.(func() uint32)
	// detailDescAppID is the schema descriptor for app_id field.
	detailDescAppID := detailFields[1].Descriptor()
	// detail.DefaultAppID holds the default value on creation for the app_id field.
	detail.DefaultAppID = detailDescAppID.Default.(func() uuid.UUID)
	// detailDescUserID is the schema descriptor for user_id field.
	detailDescUserID := detailFields[2].Descriptor()
	// detail.DefaultUserID holds the default value on creation for the user_id field.
	detail.DefaultUserID = detailDescUserID.Default.(func() uuid.UUID)
	// detailDescCoinTypeID is the schema descriptor for coin_type_id field.
	detailDescCoinTypeID := detailFields[3].Descriptor()
	// detail.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	detail.DefaultCoinTypeID = detailDescCoinTypeID.Default.(func() uuid.UUID)
	// detailDescIoType is the schema descriptor for io_type field.
	detailDescIoType := detailFields[4].Descriptor()
	// detail.DefaultIoType holds the default value on creation for the io_type field.
	detail.DefaultIoType = detailDescIoType.Default.(string)
	// detailDescIoSubType is the schema descriptor for io_sub_type field.
	detailDescIoSubType := detailFields[5].Descriptor()
	// detail.DefaultIoSubType holds the default value on creation for the io_sub_type field.
	detail.DefaultIoSubType = detailDescIoSubType.Default.(string)
	// detailDescFromCoinTypeID is the schema descriptor for from_coin_type_id field.
	detailDescFromCoinTypeID := detailFields[7].Descriptor()
	// detail.DefaultFromCoinTypeID holds the default value on creation for the from_coin_type_id field.
	detail.DefaultFromCoinTypeID = detailDescFromCoinTypeID.Default.(func() uuid.UUID)
	// detailDescIoExtra is the schema descriptor for io_extra field.
	detailDescIoExtra := detailFields[9].Descriptor()
	// detail.DefaultIoExtra holds the default value on creation for the io_extra field.
	detail.DefaultIoExtra = detailDescIoExtra.Default.(string)
	// detailDescID is the schema descriptor for id field.
	detailDescID := detailFields[0].Descriptor()
	// detail.DefaultID holds the default value on creation for the id field.
	detail.DefaultID = detailDescID.Default.(func() uuid.UUID)
	generalMixin := schema.General{}.Mixin()
	general.Policy = privacy.NewPolicies(generalMixin[0], schema.General{})
	general.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := general.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	generalMixinFields0 := generalMixin[0].Fields()
	_ = generalMixinFields0
	generalFields := schema.General{}.Fields()
	_ = generalFields
	// generalDescCreatedAt is the schema descriptor for created_at field.
	generalDescCreatedAt := generalMixinFields0[0].Descriptor()
	// general.DefaultCreatedAt holds the default value on creation for the created_at field.
	general.DefaultCreatedAt = generalDescCreatedAt.Default.(func() uint32)
	// generalDescUpdatedAt is the schema descriptor for updated_at field.
	generalDescUpdatedAt := generalMixinFields0[1].Descriptor()
	// general.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	general.DefaultUpdatedAt = generalDescUpdatedAt.Default.(func() uint32)
	// general.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	general.UpdateDefaultUpdatedAt = generalDescUpdatedAt.UpdateDefault.(func() uint32)
	// generalDescDeletedAt is the schema descriptor for deleted_at field.
	generalDescDeletedAt := generalMixinFields0[2].Descriptor()
	// general.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	general.DefaultDeletedAt = generalDescDeletedAt.Default.(func() uint32)
	// generalDescAppID is the schema descriptor for app_id field.
	generalDescAppID := generalFields[1].Descriptor()
	// general.DefaultAppID holds the default value on creation for the app_id field.
	general.DefaultAppID = generalDescAppID.Default.(func() uuid.UUID)
	// generalDescUserID is the schema descriptor for user_id field.
	generalDescUserID := generalFields[2].Descriptor()
	// general.DefaultUserID holds the default value on creation for the user_id field.
	general.DefaultUserID = generalDescUserID.Default.(func() uuid.UUID)
	// generalDescCoinTypeID is the schema descriptor for coin_type_id field.
	generalDescCoinTypeID := generalFields[3].Descriptor()
	// general.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	general.DefaultCoinTypeID = generalDescCoinTypeID.Default.(func() uuid.UUID)
	// generalDescID is the schema descriptor for id field.
	generalDescID := generalFields[0].Descriptor()
	// general.DefaultID holds the default value on creation for the id field.
	general.DefaultID = generalDescID.Default.(func() uuid.UUID)
	profitMixin := schema.Profit{}.Mixin()
	profit.Policy = privacy.NewPolicies(profitMixin[0], schema.Profit{})
	profit.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := profit.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	profitMixinFields0 := profitMixin[0].Fields()
	_ = profitMixinFields0
	profitFields := schema.Profit{}.Fields()
	_ = profitFields
	// profitDescCreatedAt is the schema descriptor for created_at field.
	profitDescCreatedAt := profitMixinFields0[0].Descriptor()
	// profit.DefaultCreatedAt holds the default value on creation for the created_at field.
	profit.DefaultCreatedAt = profitDescCreatedAt.Default.(func() uint32)
	// profitDescUpdatedAt is the schema descriptor for updated_at field.
	profitDescUpdatedAt := profitMixinFields0[1].Descriptor()
	// profit.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	profit.DefaultUpdatedAt = profitDescUpdatedAt.Default.(func() uint32)
	// profit.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	profit.UpdateDefaultUpdatedAt = profitDescUpdatedAt.UpdateDefault.(func() uint32)
	// profitDescDeletedAt is the schema descriptor for deleted_at field.
	profitDescDeletedAt := profitMixinFields0[2].Descriptor()
	// profit.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	profit.DefaultDeletedAt = profitDescDeletedAt.Default.(func() uint32)
	// profitDescAppID is the schema descriptor for app_id field.
	profitDescAppID := profitFields[1].Descriptor()
	// profit.DefaultAppID holds the default value on creation for the app_id field.
	profit.DefaultAppID = profitDescAppID.Default.(func() uuid.UUID)
	// profitDescUserID is the schema descriptor for user_id field.
	profitDescUserID := profitFields[2].Descriptor()
	// profit.DefaultUserID holds the default value on creation for the user_id field.
	profit.DefaultUserID = profitDescUserID.Default.(func() uuid.UUID)
	// profitDescCoinTypeID is the schema descriptor for coin_type_id field.
	profitDescCoinTypeID := profitFields[3].Descriptor()
	// profit.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	profit.DefaultCoinTypeID = profitDescCoinTypeID.Default.(func() uuid.UUID)
	// profitDescID is the schema descriptor for id field.
	profitDescID := profitFields[0].Descriptor()
	// profit.DefaultID holds the default value on creation for the id field.
	profit.DefaultID = profitDescID.Default.(func() uuid.UUID)
	withdrawMixin := schema.Withdraw{}.Mixin()
	withdraw.Policy = privacy.NewPolicies(withdrawMixin[0], schema.Withdraw{})
	withdraw.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := withdraw.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	withdrawMixinFields0 := withdrawMixin[0].Fields()
	_ = withdrawMixinFields0
	withdrawFields := schema.Withdraw{}.Fields()
	_ = withdrawFields
	// withdrawDescCreatedAt is the schema descriptor for created_at field.
	withdrawDescCreatedAt := withdrawMixinFields0[0].Descriptor()
	// withdraw.DefaultCreatedAt holds the default value on creation for the created_at field.
	withdraw.DefaultCreatedAt = withdrawDescCreatedAt.Default.(func() uint32)
	// withdrawDescUpdatedAt is the schema descriptor for updated_at field.
	withdrawDescUpdatedAt := withdrawMixinFields0[1].Descriptor()
	// withdraw.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	withdraw.DefaultUpdatedAt = withdrawDescUpdatedAt.Default.(func() uint32)
	// withdraw.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	withdraw.UpdateDefaultUpdatedAt = withdrawDescUpdatedAt.UpdateDefault.(func() uint32)
	// withdrawDescDeletedAt is the schema descriptor for deleted_at field.
	withdrawDescDeletedAt := withdrawMixinFields0[2].Descriptor()
	// withdraw.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	withdraw.DefaultDeletedAt = withdrawDescDeletedAt.Default.(func() uint32)
	// withdrawDescAppID is the schema descriptor for app_id field.
	withdrawDescAppID := withdrawFields[1].Descriptor()
	// withdraw.DefaultAppID holds the default value on creation for the app_id field.
	withdraw.DefaultAppID = withdrawDescAppID.Default.(func() uuid.UUID)
	// withdrawDescUserID is the schema descriptor for user_id field.
	withdrawDescUserID := withdrawFields[2].Descriptor()
	// withdraw.DefaultUserID holds the default value on creation for the user_id field.
	withdraw.DefaultUserID = withdrawDescUserID.Default.(func() uuid.UUID)
	// withdrawDescCoinTypeID is the schema descriptor for coin_type_id field.
	withdrawDescCoinTypeID := withdrawFields[3].Descriptor()
	// withdraw.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	withdraw.DefaultCoinTypeID = withdrawDescCoinTypeID.Default.(func() uuid.UUID)
	// withdrawDescAccountID is the schema descriptor for account_id field.
	withdrawDescAccountID := withdrawFields[4].Descriptor()
	// withdraw.DefaultAccountID holds the default value on creation for the account_id field.
	withdraw.DefaultAccountID = withdrawDescAccountID.Default.(func() uuid.UUID)
	// withdrawDescPlatformTransactionID is the schema descriptor for platform_transaction_id field.
	withdrawDescPlatformTransactionID := withdrawFields[5].Descriptor()
	// withdraw.DefaultPlatformTransactionID holds the default value on creation for the platform_transaction_id field.
	withdraw.DefaultPlatformTransactionID = withdrawDescPlatformTransactionID.Default.(func() uuid.UUID)
	// withdrawDescChainTransactionID is the schema descriptor for chain_transaction_id field.
	withdrawDescChainTransactionID := withdrawFields[6].Descriptor()
	// withdraw.DefaultChainTransactionID holds the default value on creation for the chain_transaction_id field.
	withdraw.DefaultChainTransactionID = withdrawDescChainTransactionID.Default.(string)
	// withdrawDescState is the schema descriptor for state field.
	withdrawDescState := withdrawFields[7].Descriptor()
	// withdraw.DefaultState holds the default value on creation for the state field.
	withdraw.DefaultState = withdrawDescState.Default.(string)
	// withdrawDescID is the schema descriptor for id field.
	withdrawDescID := withdrawFields[0].Descriptor()
	// withdraw.DefaultID holds the default value on creation for the id field.
	withdraw.DefaultID = withdrawDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.10.1"                                         // Version of ent codegen.
	Sum     = "h1:dM5h4Zk6yHGIgw4dCqVzGw3nWgpGYJiV4/kyHEF6PFo=" // Sum of ent codegen.
)
