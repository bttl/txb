package txb

import (
	"crypto/md5"
	"fmt"
	"io"
)

func computeChecksum(vk_id, app_id, app_skey string) string {
	h := md5.New()

	str := app_id + "_" + vk_id + "_" + app_skey

	io.WriteString(h, str)

	return fmt.Sprintf("%x", h.Sum(nil))
}
