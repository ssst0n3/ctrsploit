package pipe_primitive

import (
	"bytes"
	"fmt"
	"github.com/ctrsploit/ctrsploit/log"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/ioutil"
)

func MultiWrite(primitive Primitive, path string, payload []byte) (err error) {
	minOffset := primitive.MinOffset()
	if int64(len(payload)) < minOffset {
		err = fmt.Errorf("len(payload)=%d < minOffset=%d", len(payload), minOffset)
		awesome_error.CheckErr(err)
		return
	}
	content, err := ioutil.ReadFile(path)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if int64(len(content)) < minOffset {
		err = fmt.Errorf("len(content)=%d < minOffset=%d", len(content), minOffset)
		awesome_error.CheckErr(err)
		return
	}
	var i int64 = 0
	for ; i <= int64(len(content))/4096; i += 1 {
		start := minOffset + i*4096
		end := (i + 1) * 4096
		//if end > int64(len(content)) {
		//	end = int64(len(content))
		//}
		if end > int64(len(payload)) {
			end = int64(len(payload))
		}
		if start > end {
			break
		}
		if bytes.Compare(content[:start], payload[:start]) != 0 {
			err = fmt.Errorf(
				"the first %d bytes are (origin)%+v != (payload)%+v, but the rest of content will still be overwritten",
				minOffset, content[:minOffset], payload[:minOffset],
			)
			log.Logger.Warn(err)
		}
		err = primitive.Write(path, start, payload[start:end])
		if err != nil {
			return
		}
	}
	return
}
