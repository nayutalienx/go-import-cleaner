# go-import-cleaner

This utility cleans aliases in imports. The result is written to the passed file.

# example

input
```
...
import (
	"os"
	filepath "path/filepath"
	"strings"

	core "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/lookup"
	pathvar "github.com/hyperledger/fabric-sdk-go/pkg/util/pathvar"
	"github.com/spf13/cast"
)

...
```

output
```
...
import (
	"os"
	"path/filepath"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config/lookup"
	"github.com/hyperledger/fabric-sdk-go/pkg/util/pathvar"
	"github.com/spf13/cast"
)

...
```
