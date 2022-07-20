// (c) 2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"github.com/lasthyphen/beacongo/genesis"
	"github.com/lasthyphen/beacongo/ids"
	"github.com/lasthyphen/beacongo/utils/constants"
	"github.com/lasthyphen/beacongo/vms/platformvm"
)

type GenesisContainer struct {
	NetworkID       uint32
	XChainGenesisTx *platformvm.Tx
	XChainID        ids.ID
	DjtxAssetID     ids.ID
	GenesisBytes    []byte
}

func NewGenesisContainer(networkID uint32) (*GenesisContainer, error) {
	gc := &GenesisContainer{NetworkID: networkID}
	var err error
	gc.GenesisBytes, gc.DjtxAssetID, err = genesis.FromConfig(genesis.GetConfig(gc.NetworkID))
	if err != nil {
		return nil, err
	}

	gc.XChainGenesisTx, err = genesis.VMGenesis(gc.GenesisBytes, constants.AVMID)
	if err != nil {
		return nil, err
	}

	gc.XChainID = gc.XChainGenesisTx.ID()
	return gc, nil
}
