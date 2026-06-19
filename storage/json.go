package storage
import (
	"encoding/json"
	"os"
	"net/url"
	"strings"
)

func SaveJSON(site string, data any) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil { return err }
	u, _ := url.Parse(site)	
	host := strings.TrimPrefix(u.Host, "www.")
	name := strings.Split(host, ".")[0]
	filename := "data/" + name + ".json"
	return os.WriteFile(filename, jsonData, 0644)
}