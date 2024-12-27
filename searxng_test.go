package searxng

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestClient_Search(t *testing.T) {
	client, err := NewClient(&ClientOption{
		URL: "http://localhost:8080",
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range AllCategories {
		output, err := client.Search(context.Background(), &SearchInput{
			Query:      "Tokyo",
			Categories: []Category{c},
		})
		if err != nil {
			t.Fatal(err)
		}

		// category名のファイルに出力して、そこからcategory毎のgo structを作ってunion的に使
		f, err := os.Create(fmt.Sprintf("testdata/%s.json", c))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		enc := json.NewEncoder(f)
		enc.SetIndent("", "  ")
		if err := enc.Encode(output); err != nil {
			t.Fatal(err)
		}
		t.Log(c, "done")
	}
}
