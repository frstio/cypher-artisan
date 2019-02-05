package artisan

import "testing"

func TestArtisan(t *testing.T) {
	query := QueryBuilder().
		Match(
			"(w1:Wallet {address: {w1}})",
			Node("w2").Labels("Wallet").String(),
			"p = (w1)-[tx:TX*3]-(w2)",
		).
		With(
			As("p", "p"),
			As("w1", "w1"),
			As("w2", "w2"),
			As("w2.address", "recipient"),
			As(`filter(x in tx WHERE (((x.token = '0x') AND (toFloat(x.amount) = 100)) AND 
	datetime(x.last_observed).epochSeconds > -62135596800
))`, "tx2"),
		).
		Return("p", "w1", "w2", "length(p), tx2").
		Skip(5).
		Limit(10).
		Execute()

	expected := `
		MATCH 
			(w1:Wallet {address: {w1}}),
			(w2:Wallet),
			p = (w1)-[tx:TX*3]-(w2)
			
		WITH 
			p AS p, w1 AS w1, w2 AS w2, w2.address AS recipient, filter(x in tx WHERE (((x.token = '0x') AND (toFloat(x.amount) = 100)) AND 
	datetime(x.last_observed).epochSeconds > -62135596800
)) AS tx2
		RETURN 
			p, w1, w2, length(p), tx2
		SKIP	5
		LIMIT	10`

	if expected != query {
		t.Errorf("Generated query does not equal to expected:\n Expected:`%v`\n Query:`%v`", expected, query)
	}
}