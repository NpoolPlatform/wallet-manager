package profit

import (
	"fmt"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/ledger/mgr/v1/ledger/profit"

	"github.com/google/uuid"
)

func validate(info *npool.ProfitReq) error {
	if info.AppID == nil {
		logger.Sugar().Error("AppID is empty")
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Error("AppID is invalid: %v", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("AppID is invalid: %v", err))
	}

	if info.UserID == nil {
		logger.Sugar().Error("UserID is empty")
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetUserID()); err != nil {
		logger.Sugar().Error("UserID is invalid: %v", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("UserID is invalid: %v", err))
	}

	if info.CoinTypeID == nil {
		logger.Sugar().Error("CoinTypeID is empty")
		return status.Error(codes.InvalidArgument, "UserID is empty")
	}

	if _, err := uuid.Parse(info.GetCoinTypeID()); err != nil {
		logger.Sugar().Error("CoinTypeID is invalid: %v", err)
		return status.Error(codes.InvalidArgument, fmt.Sprintf("CoinTypeID is invalid: %v", err))
	}

	if info.Incoming != nil {
		incoming, err := decimal.NewFromString(info.GetIncoming())
		if err != nil {
			logger.Sugar().Error("Incoming is invalid")
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Incoming is invalid: %v", err))
		}
		if incoming.Cmp(decimal.NewFromInt(0)) < 0 {
			logger.Sugar().Error("Incoming is less than 0")
			return status.Error(codes.InvalidArgument, "Incoming is less than 0")
		}
	}

	return nil
}

func duplicate(infos []*npool.ProfitReq) error {
	keys := map[string]struct{}{}
	apps := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("Infos has invalid element %v", err))
		}

		key := fmt.Sprintf("%v:%v:%v", info.GetAppID(), info.GetUserID(), info.GetCoinTypeID())
		if _, ok := keys[key]; ok {
			return status.Error(codes.InvalidArgument, "Infos has duplicate AppID:UserID:CoinTypeID")
		}

		keys[key] = struct{}{}
		apps[info.GetAppID()] = struct{}{}
	}

	if len(apps) > 1 {
		return status.Error(codes.InvalidArgument, "Infos has different AppID")
	}

	return nil
}
