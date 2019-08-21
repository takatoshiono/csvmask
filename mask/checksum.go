package mask

import (
	"fmt"
	"hash/crc32"
)

// checksum returns CRC-32 checksum of s.
func checksum(s string) string {
	if s == "" {
		return s
	}
	return fmt.Sprintf("%08x", crc32.Checksum([]byte(s), crc32.IEEETable))
}
