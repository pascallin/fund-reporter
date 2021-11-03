package datasource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStatAPIData(t *testing.T) {
	server, err := mockServer(t)
	if err != nil {
		require.NoError(t, err)
	}
	t.Cleanup(server.Close)
	t.Run("GetStatAPIData succeed", func(t *testing.T) {
		url := fmt.Sprintf("%s/fake", server.URL)
		data, err := getStatAPIData(server.Client(), url)
		require.NoError(t, err)
		require.Equal(t, 200, data.Code)
	})
}

func mockServer(t *testing.T) (*httptest.Server, error) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		// this is where your test fixture goes
		testData := Res{
			Code: 200,
			Data: []ResData{},
			Msg:  "succeed",
		}
		err := enc.Encode(&testData)
		if err != nil {
			require.NoError(t, err)
		}
	}))
	return server, nil
}
