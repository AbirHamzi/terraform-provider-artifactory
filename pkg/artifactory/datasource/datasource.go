package datasource

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"os/exec"
)

func VerifySha256Checksum(path string, expectedSha256 string) (bool, string, string,string, error) {
	f, err := os.Open(path)
	if err != nil {
		return false,"","","", err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	hasher := sha256.New()

	if _, err := io.Copy(hasher, f); err != nil {
		return false,"","","", err
	}
	/* Debug */
	cmd1 := exec.Command("sha256sum", path)
    stdout1, err := cmd1.Output()
	if err != nil {
		return false,"","","", err
    }
	cmd2 := exec.Command("od", "--format=x1", "--read-bytes=16", path)
    stdout2, err := cmd2.Output()
	if err != nil {
		return false,"","","", err
    }
    /**********/
	return hex.EncodeToString(hasher.Sum(nil)) == expectedSha256, hex.EncodeToString(hasher.Sum(nil)),string(stdout1),string(stdout2),nil
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
