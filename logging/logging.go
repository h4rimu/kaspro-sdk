//
//  logging.go
//  kaspro-sdk
//
//  Copyright © 2019 Digiasia Bios. All rights reserved.
//

package logging

import (
	"os"

	opLogging "github.com/op/go-logging"
)

var log = opLogging.MustGetLogger("kaspro-sdk")

const INTERNAL = "internal"

func MustGetLogger(name string) *KasproSdk {
	host, err := os.Hostname()
	if err != nil {
		log.Error("", err.Error())
		host = "unknown"
	}
	ppl := &KasproSdk{opLogging.MustGetLogger(name), host}
	ppl.ExtraCalldepth = 1
	return ppl
}

type KasproSdk struct {
	*opLogging.Logger
	Hostname string
}

/*func (ppl *KasproSdk) Debug(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
	ppl.Logger.Debug("debug %s", args...)
}*/

func (ppl *KasproSdk) Debugf(userid string, string_format string, args ...interface{}) {
	ppl.Logger.Debugf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (ppl *KasproSdk) Info(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
	ppl.Logger.Info("info", args)
}

func (ppl *KasproSdk) Infof(userid string, string_format string, args ...interface{}) {
	ppl.Logger.Infof("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (ppl *KasproSdk) Error(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
	ppl.Logger.Error("err", args)
}

func (ppl *KasproSdk) Errorf(userid string, string_format string, args ...interface{}) {
	ppl.Logger.Errorf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (ppl *KasproSdk) Critical(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
	ppl.Logger.Critical("crit", args)
}

/*func (ppl *KasproSdk) Criticalf(userid string, string_format string, args ...interface{}) {
	ppl.Logger.Criticalf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
}*/

func (ppl *KasproSdk) Fatal(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
	ppl.Logger.Fatal(args)
}

func (ppl *KasproSdk) Fatalf(userid string, string_format string, args ...interface{}) {
	ppl.Logger.Fatalf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (ppl *KasproSdk) Panic(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
	ppl.Logger.Panic(args)
}

func (ppl *KasproSdk) Panicf(userid string, string_format string, args ...interface{}) {
	ppl.Logger.Panicf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (ppl *KasproSdk) Warning(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
	ppl.Logger.Warning("warning", args)
}

func (ppl *KasproSdk) Warningf(userid string, string_format string, args ...interface{}) {
	ppl.Logger.Warningf("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
}

func (ppl *KasproSdk) Notice(userid string, args ...interface{}) {
	args = append([]interface{}{"[" + ppl.Hostname + "] [" + userid + "]"}, args...)
	ppl.Logger.Notice("notice", args)
}

func (ppl *KasproSdk) Noticef(userid string, string_format string, args ...interface{}) {
	ppl.Logger.Noticef("["+ppl.Hostname+"] ["+userid+"] "+string_format, args...)
}
