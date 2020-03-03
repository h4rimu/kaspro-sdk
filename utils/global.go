package utils

import (
	"github.com/h4rimu/kaspro-sdk/app/types"
	"github.com/h4rimu/kaspro-sdk/database"
	opLogging "github.com/h4rimu/kaspro-sdk/logging"
)

var MMEN *types.MessageMap
var MMID *types.MessageMap
var DBModel *database.DBModel

var log = opLogging.MustGetLogger("kaspro-sdk")
