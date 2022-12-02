package typedeclimport

import subpkg "github.com/alkiranet/gogo-protobuf/test/typedeclimport/subpkg"

type SomeMessage struct {
	Imported subpkg.AnotherMessage
}
