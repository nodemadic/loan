package keeper

import (
	"github.com/nodemadic/loan/x/loan/types"
)

var _ types.QueryServer = Keeper{}
