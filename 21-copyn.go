package main

import (
	"fmt"
	"io"
)

//START OMIT

func (c *Client) CopyWithQuota(copied int64, w io.Writer, r io.Reader, blockSize int64, fileID string, force bool) (int64, error) {
	for {
		n, err := io.CopyN(w, r, blockSize)
		copied += n
		// Check quota if copied anything
		if n > 0 {
			status, _, errUpdate := c.UpdateUsage(fileID, copied, force)
			//END OMIT
			if errUpdate != nil {
				return copied, fmt.Errorf("update usage error: %s", err)
			}
			if status != OK {
				if status == EDQUOT {
					return copied, ErrQuotaExceeded
				}
				return copied, fmt.Errorf("quota check failed: %v", status)
			}
		}
		// Stop if reached EOF
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}

		// If error is not EOF - fail
		if err != nil {
			return copied, fmt.Errorf("error during copy %d bytes: %s", blockSize, err)
		}
	}
	return copied, nil
}

