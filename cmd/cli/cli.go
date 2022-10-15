package main

import (
	"encoding/json"
	"fmt"
	"github.com/yu-org/yu/common"
	"github.com/yu-org/yu/core/keypair"
	"github.com/yu-org/yu/example/client/callchain"
	"math/big"
	"os"
	"time"
	"uask-chain/types"
)

func main() {
	pub, priv := keypair.GenSrKeyWithSecret([]byte("uask-chain"))

	action := os.Args[1]
	titleOrId := os.Args[2]
	content := os.Args[3]

	fmt.Printf("%s %s %s \n", action, titleOrId, content)

	var (
		tripod string
		exec   string
		params []byte

		err error
	)
	switch action {
	case "ask":
		info := &types.QuestionAddRequest{
			Title: titleOrId,
			Content: &types.StoreInfo{
				OnchainStore: true,
				Content:      []byte(content),
			},
			Tags:         nil,
			TotalRewards: big.NewInt(100),
			Timestamp:    time.Now().String(),
			Recommender:  common.Address{},
		}
		params, err = json.Marshal(info)
		if err != nil {
			fmt.Println("marshal ask err: ", err)
			os.Exit(1)
		}
		tripod = "question"
		exec = "AddQuestion"
	case "answer":
		info := &types.AnswerAddRequest{
			QID: titleOrId,
			Content: &types.StoreInfo{
				OnchainStore: true,
				Content:      []byte(content),
			},
			Timestamp:   time.Now().String(),
			Recommender: common.Address{},
		}
		params, err = json.Marshal(info)
		if err != nil {
			fmt.Println("marshal answer err: ", err)
			os.Exit(1)
		}
		tripod = "answer"
		exec = "AddAnswer"
	case "comment":
		info := &types.CommentAddRequest{
			AID: titleOrId,
			CID: titleOrId,
			Content: &types.StoreInfo{
				OnchainStore: true,
				Content:      []byte(content),
			},
			Timestamp: time.Now().String(),
		}
		params, err = json.Marshal(info)
		if err != nil {
			fmt.Println("marshal comment err: ", err)
			os.Exit(1)
		}
		tripod = "comment"
		exec = "AddComment"
	}

	callchain.CallChainByExec(callchain.Websocket, priv, pub, &common.Ecall{
		TripodName: tripod,
		ExecName:   exec,
		Params:     string(params),
		LeiPrice:   0,
	})
}
