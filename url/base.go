package urls

import (
	"fmt"
)

func Base(name string) string {
	baseName := fmt.Sprintf("https://www.google.com/search?q=%s", name)
	return baseName
}
