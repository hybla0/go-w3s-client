package w3s

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ipfs/go-cid"
	w3http "github.com/web3-storage/go-w3s-client/http"
)

func (c *client) Get(ctx context.Context, cid cid.Cid) (*w3http.Web3Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/car/%s", c.cfg.endpoint, cid), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.cfg.token))
	req.Header.Add("X-Client", "web3.storage/go")
	res, err := c.hc.Do(req)
	return w3http.NewWeb3Response(res, c.bsvc), err
}