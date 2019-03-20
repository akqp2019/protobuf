package typedeclimport

import subpkg "github.com/akqp2019/protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
