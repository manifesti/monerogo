// Copyright 2017 Marin Basic <marin@marin-basic.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package monerogo

type WalletClient struct {
	endpoint string
}

// NewWalletClient - Creates new wallet client
func NewWalletClient(endpoint string) *WalletClient {
	return &WalletClient{endpoint: endpoint}
}

// StopWallet -  Stops the Wallet.
func (c *WalletClient) StopWallet() error {
	if err := call(c.endpoint, "stop_wallet", nil, nil); err != nil {
		return err
	}
	return nil
}
func (c *WalletClient) GetAddress() (string, error) {
	var address string
	if err := call(c.endpoint, "getaddress", nil, address); err != nil {
		return "", err
	}
	return address, nil
}

// RescanWalletChain - Rescans the blockchain for the wallet in use.
func (c *WalletClient) RescanWalletChain() error {
	if err := call(c.endpoint, "rescan_blockchain", nil, nil); err != nil {
		return err
	}
	return nil
}

// SignString - Signs given string with wallet.
func (c *WalletClient) SignString(data string) (string, error) {
	var signedString string
	if err := call(c.endpoint, "sign", nil, &signedString); err != nil {
		return signedString, err
	}
	return signedString, nil
}

// OpenWallet - Opens the wallet from the wallet directory given on server start flags.
func (c *WalletClient) OpenWallet(filename string, password string) error {
	req := struct {
		fileName string `json:"filename"`
		passWord string `json:"password"`
	}{filename, password}
	if err := call(c.endpoint, "open_wallet", req, nil); err != nil {
		return err
	}
	return nil
}

// CreateWallet - Creates a wallet to the wallet directory.
func (c *WalletClient) CreateWallet(filename string, password string, language string) error {
	req := struct {
		fileName string `json:"filename"`
		passWord string `json:"password"`
		language string `json:"language"`
	}{filename, password, language}
	if err := call(c.endpoint, "create_wallet", req, nil); err != nil {
		return err
	}
	return nil
}

// StartMining - Starts mining with the connected Daemon.
func (c *WalletClient) StartMining(threadcount uint, background bool, battery bool) error {
	req := struct {
		threadCount uint `json:"threads_count"`
		backGround  bool `json:"do_background_mining"`
		battery     bool `json:"ignore_battery"`
	}{threadcount, background, battery}
	if err := call(c.endpoint, "start_mining", req, nil); err != nil {
		return err
	}
	return nil
}

// StopMining - Stops mining with the daemon.
func (c *WalletClient) StopMining() error {
	if err := call(c.endpoint, "stop_mining", nil, nil); err != nil {
		return err
	}
	return nil
}

// IncomingTransfers - Return a list of incoming transfers to the wallet.
func (c *WalletClient) IncomingTransfers(transfertype string) (IncomingTransferResponse, error) {
	req := struct {
		transferType string `json:"transfer_type"`
	}{transfertype}

	var resp IncomingTransferResponse
	if err := call(c.endpoint, "incoming_transfers", req, resp); err != nil {
		return resp, err
	}
	return resp, nil
}
