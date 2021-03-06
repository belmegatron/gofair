package streaming

import (
	"github.com/sirupsen/logrus"

	"github.com/belmegatron/gofair/streaming/models"
)

type IStream interface {
	OnSubscribe(ChangeMessage models.MarketChangeMessage)
	OnResubscribe(ChangeMessage models.MarketChangeMessage)
	OnHeartbeat(ChangeMessage models.MarketChangeMessage)
	OnUpdate(ChangeMessage models.MarketChangeMessage)
}

type MarketStream struct {
	log           *logrus.Logger
	OutputChannel chan MarketBook
	Cache         map[string]MarketCache
	InitialClk    string
	Clk           string
}

func getMarketIDs(mcm models.MarketChangeMessage) []string {
	marketIDs := make([]string, 0)
	for _, marketChange := range mcm.Mc {
		marketIDs = append(marketIDs, marketChange.ID)
	}
	return marketIDs
}

func (ms *MarketStream) OnSubscribe(changeMessage models.MarketChangeMessage) {

	marketIDs := getMarketIDs(changeMessage)

	ms.log.WithFields(logrus.Fields{
		"marketIDs": marketIDs,
	}).Debug("Subscribed to Betfair Market Changes")

	for _, marketChange := range changeMessage.Mc {

		marketCache := CreateMarketCache(&changeMessage, marketChange)
		ms.Cache[marketChange.ID] = *marketCache
	}

}

func (ms *MarketStream) OnResubscribe(changeMessage models.MarketChangeMessage) {
	marketIDs := getMarketIDs(changeMessage)
	ms.log.WithFields(logrus.Fields{
		"marketIDs": marketIDs,
	}).Debug("Resubscribed to Betfair Market Changes")
}

func (ms *MarketStream) OnHeartbeat(changeMessage models.MarketChangeMessage) {
	ms.log.Debug("Heartbeat")
}

func (ms *MarketStream) OnUpdate(changeMessage models.MarketChangeMessage) {

	if ms.InitialClk == "" {
		ms.InitialClk = changeMessage.Clk
	}

	ms.Clk = changeMessage.Clk

	for _, marketChange := range changeMessage.Mc {

		if marketCache, ok := ms.Cache[marketChange.ID]; ok {
			marketCache.UpdateCache(&changeMessage, marketChange)
			ms.OutputChannel <- marketCache.Snap()
		} else {
			marketCache := CreateMarketCache(&changeMessage, marketChange)
			ms.Cache[marketChange.ID] = *marketCache
			ms.OutputChannel <- marketCache.Snap()
		}
	}
}
