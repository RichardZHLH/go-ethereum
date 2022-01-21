// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package keystore

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

// add by Jacob begin
// GetWanAddress represents the wallet to retrieve corresponding wanchain public address for a specific ordinary account/address
func (w *keystoreWallet) GetWanAddress(account accounts.Account) (common.WAddress, error) {
	// Make sure the requested account is contained within
	if account.Address != w.account.Address {
		return common.WAddress{}, accounts.ErrUnknownAccount
	}
	if account.URL != (accounts.URL{}) && account.URL != w.account.URL {
		return common.WAddress{}, accounts.ErrUnknownAccount
	}
	// Account seems valid, request the keystore to retrieve
	return w.keystore.GetWanAddress(account)
}

// SignHashWithPassphrase implements accounts.Wallet, attempting to sign the
// given hash with the given account using passphrase as extra authentication.
func (w *keystoreWallet) SignHashWithPassphrase(account accounts.Account, passphrase string, hash []byte) ([]byte, error) {
	// Make sure the requested account is contained within
	if account.Address != w.account.Address {
		return nil, accounts.ErrUnknownAccount
	}
	if account.URL != (accounts.URL{}) && account.URL != w.account.URL {
		return nil, accounts.ErrUnknownAccount
	}
	// Account seems valid, request the keystore to sign
	return w.keystore.SignHashWithPassphrase(account, passphrase, hash)
}

func (w *keystoreWallet) ComputeOTAPPKeys(account accounts.Account, AX, AY, BX, BY string) ([]string, error) {
	// Make sure the requested account is contained within
	if account.Address != w.account.Address {
		return nil, accounts.ErrUnknownAccount
	}
	if account.URL != (accounts.URL{}) && account.URL != w.account.URL {
		return nil, accounts.ErrUnknownAccount
	}

	// Account seems valid, request the keystore to process
	return w.keystore.ComputeOTAPPKeys(account, AX, AY, BX, BY)
}

// add by Jacob end
