package utils

import "fmt"

func FormatAllNetAddr(port uint) string {
	return fmt.Sprintf(":%d", port)
}
