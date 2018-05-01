package util

import (
	"bufio"
	"os"
)

func GenKey(hashTag string, primaryKey string) string {
	ret := "{" + hashTag + "}." + primaryKey
	return ret
}

func GetAddrs(configFile string) []string {
	file, err := os.Open(configFile)
	if err != nil {
		return nil
	}
	defer file.Close()

	ret := make([]string, 1)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}
	return ret
}
