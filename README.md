# go-import-cleaner

A command-line utility for cleaning and organizing Go import statements by removing unnecessary aliases and sorting imports alphabetically.

## Features

- **Remove import aliases**: Automatically removes custom aliases from import statements
- **Sort imports**: Organizes imports alphabetically for better readability
- **Deduplicate imports**: Removes duplicate import statements
- **In-place editing**: Modifies files directly, preserving the original structure
- **Multiple import formats**: Supports both single-line and multi-line import blocks

## Installation

```bash
go get github.com/nayutalienx/go-import-cleaner
```

Or build from source:

```bash
git clone https://github.com/nayutalienx/go-import-cleaner.git
cd go-import-cleaner
go build -o go-import-cleaner main.go
```

## Usage

```bash
go-import-cleaner <path-to-go-file>
```

### Examples

**Basic usage:**
```bash
go-import-cleaner main.go
```

**Get help:**
```bash
go-import-cleaner help
# or
go-import-cleaner --help
# or
go-import-cleaner -help
```

### Before and After Example

**Input file:**
```go
package main

import (
	"os"
	filepath "path/filepath"
	"strings"

	core "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/lookup"
	pathvar "github.com/hyperledger/fabric-sdk-go/pkg/util/pathvar"
	"github.com/spf13/cast"
)

func main() {
	// Your code here
}
```

**Output file (after processing):**
```go
package main

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/lookup"
	"github.com/hyperledger/fabric-sdk-go/pkg/util/pathvar"
	"github.com/spf13/cast"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Your code here
}
```

## Supported Import Formats

The tool supports cleaning both single-line and multi-line import statements:

**Multi-line imports (most common):**
```go
import (
	alias "package/path"
	"another/package"
)
```

**Single-line aliased imports:**
```go
import alias "package/path"
```

## Limitations

- Only processes `.go` files
- Skips files with single imports without aliases
- Modifies files in-place (creates a backup if needed)
- Requires valid Go file structure

## How It Works

1. **Validation**: Ensures the input file has a `.go` extension
2. **Parsing**: Reads and analyzes import statements in the file
3. **Cleaning**: Removes aliases while preserving the actual package paths
4. **Sorting**: Alphabetically sorts all imports
5. **Deduplication**: Removes any duplicate import entries
6. **Writing**: Saves the cleaned imports back to the original file

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is open source and available under the [MIT License](LICENSE).

## Development

To run the tool locally:

```bash
go run main.go <your-file.go>
```

To build the executable:

```bash
go build -o go-import-cleaner main.go
```
