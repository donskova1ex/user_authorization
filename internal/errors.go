package internal

import "errors"

var QueryError = errors.New("db query error")
var ErrorRowScan = errors.New("db row scan error")
