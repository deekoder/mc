// --------  DO NOT EDIT --------
// this is an autogenerated file

package main

import (
	"net/http"
	"time"
)

// Version autogenerated
var Version = "2015-06-18T01:29:09.838113535Z"

// Tag is of following format
//
//   [[STRING]-[EPOCH]
//
// STRING is release string of your choice.
// EPOCH is unix seconds since Jan 1, 1970 UTC.
var Tag = "release-1434590949"

// getVersion -
func getVersion() string {
	t, _ := time.Parse(time.RFC3339Nano, Version)
	if t.IsZero() {
		return ""
	}
	return t.Format(http.TimeFormat)
}