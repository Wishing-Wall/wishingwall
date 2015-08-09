package protocol

import (
	"errors"
	"fmt"
	"quickstart/conf"
	"quickstart/models/bitcoinchain"
	"quickstart/models/dbutil"
	"quickstart/models/wallet"
	"time"
)

func SendMessage(send conf.DB_send) error {
	MsgTx, _, err := wallet.CreateRawTransaction(send.RelayAddr, send.Message)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to createrawtransaction %v\r\n", err))
	}
	hash, err := wallet.SendRawTransaction(MsgTx)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to sendrawtransaction %v\r\n", err))
	}
	send.Tx_hash = hash.String()
	send.IsSent = true
	err = dbutil.UpdateSend(send)
	return err
}

func SendLoop() {
	var LastBlockIndex uint64 = 0
	for {
		time.Sleep(1 * time.Second)
		BlockIndex, err := bitcoinchain.GetBlockCount()
		if err != nil {
			fmt.Printf("Get block count failed continue:%v\r\n", err)
			continue
		}
		if BlockIndex <= LastBlockIndex {
			fmt.Printf("blockindex[%v] < LastBlockIndex[%d]\r\n", BlockIndex, LastBlockIndex)
			continue
		}
		UnspentList, err := wallet.UnspentList(3)
		if err != nil {
			fmt.Printf("Failed to get unspendlist %v\r\n", err)
			continue
		}
		UnSendList, err := dbutil.GetAllUnsentMessage(24 * time.Hour)
		if err != nil {
			fmt.Printf("Failed to get unsendlist %v\r\n", err)
			continue
		}
		//fmt.Printf("found unspentlist %v\r\n ", UnspentList)
		//fmt.Printf("found unsendlist %v\r\n", UnSendList)
		for _, spent := range UnspentList {
			for _, send := range UnSendList {
				//fmt.Printf("spent addr %v, send addr %v\r\n", spent.Address, send.RelayAddr)

				if spent.Address == send.RelayAddr {

					amount := uint64(spent.Amount * float64(conf.COIN))
					//fmt.Printf("amount is %v, Minmoney is %v\r\n", amount, send.MinMoney)
					if amount >= send.MinMoney {
						err = SendMessage(send)
						if err != nil {
							fmt.Printf("faield to send message %v\r\n", err)
						}
					}
				}
			}
		}
	}

}
