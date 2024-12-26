package internal

import "errors"

var QueryError = errors.New("db query error")
var ErrorRowScan = errors.New("db row scan error")
var ErrReadRows = errors.New("can not read rows")
var ErrNotFound = errors.New("entity not found")
var ErrNotDelete = errors.New("entity not delete")
