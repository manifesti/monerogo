// Copyright 2017 manifesti. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package monerogo

// BalanceResponse
// balance - unsigned int; The total balance of the current monero-wallet-rpc in session.
// unlocked_balance - unsigned int; Unlocked funds are those funds that are sufficiently deep enough in the Monero blockchain to be considered safe to spend.
// status - string; General RPC error code. "OK" means everything looks good.
type BalanceResponse struct {
	Lockedbalance   uint `json:"balance"`
	Unlockedbalance uint `json:"unlocked_balance"`
}

// AddressResponse
// address - string; The 95-character hex address string of the monero-wallet-rpc in session.
// status - string; General RPC error code. "OK" means everything looks good.
type AddressResponse struct {
	Address string `json:"address"`
}

// WalletHeight
// height - unsigned int; The current monero-wallet-rpc's blockchain height. If the wallet has been offline for a long time, it may need to catch up with the daemon.
// status - string; General RPC error code. "OK" means everything looks good.
type WalletHeight struct {
	Height uint `json:"height"`
}

// TransferDestinations
// amount - unsigned int; Amount to send to each destination, in atomic units.
// address - string; Destination public address.
type TransferDestinations struct {
	Amount  uint   `json:"amount"`
	Address string `json:"address"`
}

// TransferRequest
// destinations - array of destinations (TransferDestinations) to receive XMR
// mixin - unsigned int; Number of outpouts from the blockchain to mix with (0 means no mixing).
// unlock_time - unsigned int; Number of blocks before the monero can be spent (0 to not add a lock).
// payment_id - string; (Optional) Random 32-byte/64-character hex string to identify a transaction.
// get_tx_key - boolean; (Optional) Return the transaction key after sending.
// Priority - unsigned int; Set a priority for the transaction. Accepted Values are: 0-3 for: default, unimportant, normal, elevated, priority.
// do_not_relay - boolean; (Optional) If true, the newly created transaction will not be relayed to the monero network. (Defaults to false)
// get_tx_hex - boolean; Return the transaction as hex string after sending
type TransferRequest struct {
	Destinations []TransferDestinations `json:"destinations"`
	Mixin        uint                   `json:"mixin"`
	UnlockTime   uint                   `json:"unlock_time"`
	PaymentId    string                 `json:"payment_id,omitempty"`
	GetTxKey     bool                   `json:"get_tx_key,omitempty"`
	Priority     uint                   `json:"priority"`
	DoNotRelay   bool                   `json:"do_not_relay,omitempty"`
	GetTxHex     bool                   `json:"get_tx_hex"`
}

// TransferResponse
// Fee - Integer value of the fee charged for the txn.
// tx_hash - String for the publically searchable transaction hash
// tx_key - String for the transaction key if get_tx_key is true, otherwise, blank string.
// tx_blob - Transaction as hex string if get_tx_hex is true
type TransferResponse struct {
	Fee    uint   `json:"fee"`
	TxHash string `json:"tx_hash"`
	TxKey  string `json:"tx_key"`
	TxBlob string `json:"tx_blob, omitempty"`
}

// SweepDustResponse
// tx_hash_list - list of: string
// status - string; General RPC error code. "OK" means everything looks good.
type SweepDustResponse struct {
	TxHashList []string `json:"tx_hash_list"`
}

// SweepAllRequest
// address - string; Destination public address.
// priority - unsigned int; (Optional)
// mixin - unsigned int; Number of outpouts from the blockchain to mix with (0 means no mixing).
// unlock_time - unsigned int; Number of blocks before the monero can be spent (0 to not add a lock).
// payment_id - string; (Optional) Random 32-byte/64-character hex string to identify a transaction.
// get_tx_keys - boolean; (Optional) Return the transaction keys after sending.
// below_amount - unsigned int; (Optional)
// do_not_relay - boolean; (Optional)
// get_tx_hex - boolean; (Optional) return the transactions as hex encoded string.
type SweepAllRequest struct {
	Address     string `json:"address"`
	Priority    uint   `json:"priority, omitempty"`
	Mixin       uint   `json:"mixin"`
	UnlockTime  uint   `json:"unlock_time"`
	PaymentId   string `json:"payment_id, omitempty"`
	GetTxKeys   bool   `json:"get_tx_keys, omitempty"`
	BelowAmount uint   `json:"below_amount, omitempty"`
	DoNotRelay  bool   `json:"do_not_relay, omitempty"`
	GetTxHex    bool   `json:"get_tx_hex, omitempty"`
}

// SweepAllResponse
// tx_hash_list - array of string;
// tx_key_list - array of string;
// tx_blob_list - array of string;
type SweepAllResponse struct {
	TxHashList []string `json:"tx_hash_list"`
	TxKeyList  []string `json:"tx_key_list"`
	TxBlobList []string `json:"tx_blob_list, omitempty"`
}

// Payments
// payment_id - string
// tx_hash - string
// amount - unsigned int
// block_height - unsigned int
// unlock_time - unsigned int
type Payments struct {
	PaymentId   string `json:"payment_id"`
	TxHash      string `json:"tx_hash"`
	Amount      uint   `json:"amount"`
	BlockHeight uint   `json:"block_height"`
	UnlockTime  uint   `json:"unlock_time"`
}

// GetPaymentsResponse
// Payments - array of Payments
// status - string; General RPC error code. "OK" means everything looks good.
type GetPaymentsResponse struct {
	Payments []Payments
}

// Transfer
// txid - string;
// payment_id - string;
// height - unsigned int;
// timestamp - unsigned int;
// amount - unsigned int;
// fee - unsigned int;
// note - string;
// destinations - std::list;
// type - string;
type Transfer struct {
	TxId         string                 `json:"txid"`
	PaymentId    string                 `json:"payment_id"`
	Height       uint                   `json:"height"`
	Timestamp    uint                   `json:"timestamp"`
	Amount       uint                   `json:"amount"`
	Fee          uint                   `json:"fee"`
	Note         string                 `json:"note"`
	Destinations []TransferDestinations `json:"destinations, omitempty"`
	Type         string                 `json:"type"`
}

// GetTransfersRequest
// in - boolean;
// out - boolean;
// pending - boolean;
// failed - boolean;
// pool - boolean;
// filter_by_height - boolean;
// min_height - unsigned int;
// max_height - unsigned int;
type GetTransfersRequest struct {
	In             bool `json:"in"`
	Out            bool `json:"out"`
	Pending        bool `json:"pending"`
	Failed         bool `json:"failed"`
	Pool           bool `json:"pool"`
	FilterByHeight bool `json:"filter_by_height"`
	MaxHeight      uint `json:"max_height"`
	MinHeight      uint `json:"min_height"`
}

// GetTransfersResponse
// in - array of Transfers
// out - array of Transfers
// pending - array of Transfers
// failed - array of Transfers
// pool - array of Transfers
type GetTransfersResponse struct {
	In      []Transfer `json:"in"`
	Out     []Transfer `json:"out"`
	Pending []Transfer `json:"pending"`
	Failed  []Transfer `json:"failed"`
	Pool    []Transfer `json:"pool"`
}

// GetTransferByIDResponse - maybe useless if use only Transfer struct?
// txid - string;
// payment_id - string;
// height - unsigned int;
// timestamp - unsigned int;
// amount - unsigned int;
// fee - unsigned int;
// note - string;
// type - string;
type GetTransferByIDResponse struct {
	Amount    uint   `json:"amount"`
	Fee       uint   `json:"fee"`
	Height    uint   `json:"height"`
	Note      string `json:"note"`
	PaymentID string `json:"payment_id"`
	Timestamp uint   `json:"timestamp"`
	TxID      string `json:"txid"`
	Type      string `json:"type"`
}

// IncomingTransfer
// amount - unsigned int
// spent - boolean
// global_index - unsigned int; Mostly internal use, can be ignored by most users.
// tx_hash - string; Several incoming transfers may share the same hash if they were in the same transaction.
// tx_size - unsigned int
type IncomingTransfer struct {
	Amount      uint   `json:"amount"`
	Spent       bool   `json:"spent"`
	GlobalIndex uint   `json:"global_index"`
	TxHash      string `json:"tx_hash"`
	TxSize      uint   `json:"tx_size"`
}

// IncomingTransferResponse - response containing an array of transfers and the RPC status.
type IncomingTransferResponse struct {
	Transfers []IncomingTransfer `json:"transfers"`
}
