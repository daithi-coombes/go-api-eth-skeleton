package bot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestHandler(t *testing.T) {

	res := httptest.NewRecorder()

	payload := Update{}
	payload.Message.Text = "/tec dao get proposals"
	reqBody, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "URL", bytes.NewBuffer(reqBody))

	Handler(res, req)

	resp := res.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	spew.Dump(body)
	spew.Dump(string(body))
}
