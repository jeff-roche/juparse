# juparse
junit test output parser in go

## Install
```bash
$ go install github.com/jeff-roche/juparse@latest
```

## Usage
```bash
$ juparse -file MyTestOutput.xml
```

### Flags
- `-v` Verbosity - Show reason for failed and skipped tests
- `-passed` Filter to passed tests
- `-skipped` Filter to skipped tests
- `-failed` Filter to failed tests

**Filter flags can be combined to create specific outputs:**
    
```bash
# This will show failed and skipped tests with the reason
$ juparse -file MyTestOutput.xml -v -skipped -failed
```