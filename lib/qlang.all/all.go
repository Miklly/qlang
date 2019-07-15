package qall

import (
	"github.com/miklly/qlang/lib/bufio"
	"github.com/miklly/qlang/lib/bytes"
	"github.com/miklly/qlang/lib/crypto/md5"
	"github.com/miklly/qlang/lib/encoding/hex"
	"github.com/miklly/qlang/lib/encoding/json"
	"github.com/miklly/qlang/lib/eqlang"
	"github.com/miklly/qlang/lib/errors"
	"github.com/miklly/qlang/lib/io"
	"github.com/miklly/qlang/lib/io/ioutil"
	"github.com/miklly/qlang/lib/math"
	"github.com/miklly/qlang/lib/meta"
	"github.com/miklly/qlang/lib/net/http"
	"github.com/miklly/qlang/lib/os"
	"github.com/miklly/qlang/lib/path"
	"github.com/miklly/qlang/lib/reflect"
	"github.com/miklly/qlang/lib/runtime"
	"github.com/miklly/qlang/lib/strconv"
	"github.com/miklly/qlang/lib/strings"
	"github.com/miklly/qlang/lib/sync"
	"github.com/miklly/qlang/lib/terminal"
	"github.com/miklly/qlang/lib/tpl/extractor"
	"github.com/miklly/qlang/lib/version"
	qlang "github.com/miklly/qlang/spec"

	// qlang builtin modules
	_ "github.com/miklly/qlang/lib/builtin"
	_ "github.com/miklly/qlang/lib/chan"
)

// -----------------------------------------------------------------------------

// Copyright prints qlang copyright information.
//
func Copyright() {
	version.Copyright()
}

// InitSafe inits qlang and imports modules.
//
func InitSafe(safeMode bool) {

	qlang.SafeMode = safeMode

	qlang.Import("", math.Exports) // import math as builtin package
	qlang.Import("", meta.Exports) // import meta package
	qlang.Import("bufio", bufio.Exports)
	qlang.Import("bytes", bytes.Exports)
	qlang.Import("md5", md5.Exports)
	qlang.Import("io", io.Exports)
	qlang.Import("ioutil", ioutil.Exports)
	qlang.Import("hex", hex.Exports)
	qlang.Import("json", json.Exports)
	qlang.Import("errors", errors.Exports)
	qlang.Import("eqlang", eqlang.Exports)
	qlang.Import("math", math.Exports)
	qlang.Import("os", os.Exports)
	qlang.Import("", os.InlineExports)
	qlang.Import("path", path.Exports)
	qlang.Import("http", http.Exports)
	qlang.Import("reflect", reflect.Exports)
	qlang.Import("runtime", runtime.Exports)
	qlang.Import("strconv", strconv.Exports)
	qlang.Import("strings", strings.Exports)
	qlang.Import("sync", sync.Exports)
	qlang.Import("terminal", terminal.Exports)
	qlang.Import("extractor", extractor.Exports)
}

// -----------------------------------------------------------------------------
