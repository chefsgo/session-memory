package session_buntdb

import (
	"github.com/chefsgo/chef"
	sb "github.com/chefsgo/session-buntdb"
)

func init() {
	chef.Register("memory", sb.Driver(":memory:"))
}
