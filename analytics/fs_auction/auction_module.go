package fs_auction

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/chasex/glog"
	"github.com/prebid/prebid-server/analytics"
)

type RequestType string

const (
	AUCTION RequestType = "/openrtb2/auction"
)

//Module that can perform transactional logging
type AuctionFileLogger struct {
	Logger *glog.Logger
}

//Writes AuctionObject to file
func (f *AuctionFileLogger) LogAuctionObject(ao *analytics.AuctionObject) {
	//Code to parse the object and log in a way required
	var b bytes.Buffer
	b.WriteString(jsonifyAuctionObject(ao))
	f.Logger.Debug(b.String())
	f.Logger.Flush()
}

//Writes VideoObject to file
func (f *AuctionFileLogger) LogVideoObject(vo *analytics.VideoObject) {
	// do nothing
}

//Logs SetUIDObject to file
func (f *AuctionFileLogger) LogSetUIDObject(so *analytics.SetUIDObject) {
	// do nothing
}

//Logs CookieSyncObject to file
func (f *AuctionFileLogger) LogCookieSyncObject(cso *analytics.CookieSyncObject) {
	// do nothing
}

//Logs AmpObject to file
func (f *AuctionFileLogger) LogAmpObject(ao *analytics.AmpObject) {
	// do nothing
}

//Method to initialize the analytic module
func NewAuctionFileLogger(filename string) (analytics.PBSAnalyticsModule, error) {
	options := glog.LogOptions{
		File:  filename,
		Flag:  glog.LstdNull,
		Level: glog.Ldebug,
	}
	if logger, err := glog.New(options); err == nil {
		return &AuctionFileLogger{
			logger,
		}, nil
	} else {
		return nil, err
	}
}

type fileAuctionObject analytics.AuctionObject

func jsonifyAuctionObject(ao *analytics.AuctionObject) string {
	type alias analytics.AuctionObject
	b, err := json.Marshal(&struct {
		Type RequestType `json:"type"`
		*alias
	}{
		Type:  AUCTION,
		alias: (*alias)(ao),
	})

	if err == nil {
		return string(b)
	} else {
		return fmt.Sprintf("Transactional Logs Error: Auction object badly formed %v", err)
	}
}
