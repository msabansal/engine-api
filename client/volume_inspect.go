package client

import (
	"encoding/json"
	"net/http"

	"github.com/docker/engine-api/types"
)

// VolumeInspect returns the information about a specific volume in the docker host.
func (cli *Client) VolumeInspect(volumeID string) (types.Volume, error) {
	var volume types.Volume
	resp, err := cli.get("/volumes/"+volumeID, nil, nil)
	if err != nil {
		if resp.statusCode == http.StatusNotFound {
			return volume, volumeNotFoundError{volumeID}
		}
		return volume, err
	}
	err = json.NewDecoder(resp.body).Decode(&volume)
	ensureReaderClosed(resp)
	return volume, err
}
