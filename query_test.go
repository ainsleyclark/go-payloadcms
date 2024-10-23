package payloadcms

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	t.Parallel()

	t.Run("Equals", func(t *testing.T) {
		t.Parallel()
		qb := Query().Equals("field", "value")
		assert.Equal(t, "where%5Bfield%5D%5Bequals%5D=value", qb.Build())
	})

	t.Run("NotEquals", func(t *testing.T) {
		t.Parallel()
		qb := Query().NotEquals("field", "value")
		assert.Equal(t, "where%5Bfield%5D%5Bnot_equals%5D=value", qb.Build())
	})

	t.Run("GreaterThan", func(t *testing.T) {
		t.Parallel()
		qb := Query().GreaterThan("field", "10")
		assert.Equal(t, "where%5Bfield%5D%5Bgreater_than%5D=10", qb.Build())
	})

	t.Run("LessThan", func(t *testing.T) {
		t.Parallel()
		qb := Query().LessThan("field", "5")
		assert.Equal(t, "where%5Bfield%5D%5Bless_than%5D=5", qb.Build())
	})

	t.Run("In", func(t *testing.T) {
		t.Parallel()
		qb := Query().In("field", []string{"val1", "val2", "val3"})
		assert.Equal(t, "where%5Bfield%5D%5Bin%5D=val1%2Cval2%2Cval3", qb.Build())
	})

	t.Run("And", func(t *testing.T) {
		t.Parallel()
		subQuery := Query().Equals("fieldA", "valueA").GreaterThan("fieldB", "20")
		qb := Query().And(subQuery)

		expected := url.Values{}
		expected.Add("where[and][][where[fieldA][equals]]", "valueA")
		expected.Add("where[and][][where[fieldB][greater_than]]", "20")

		assert.Equal(t, expected.Encode(), qb.Build())
	})

	t.Run("Or", func(t *testing.T) {
		t.Parallel()
		subQuery := Query().LessThan("fieldC", "15").Exists("fieldD", true)
		qb := Query().Or(subQuery)

		expected := url.Values{}
		expected.Add("where[or][][where[fieldC][less_than]]", "15")
		expected.Add("where[or][][where[fieldD][exists]]", "true")

		assert.Equal(t, expected.Encode(), qb.Build())
	})

	t.Run("Exists", func(t *testing.T) {
		t.Parallel()
		qb := Query().Exists("field", true)
		assert.Equal(t, "where%5Bfield%5D%5Bexists%5D=true", qb.Build())

		qb = Query().Exists("field", false)
		assert.Equal(t, "where%5Bfield%5D%5Bexists%5D=false", qb.Build())
	})

	t.Run("Build_Empty", func(t *testing.T) {
		t.Parallel()
		qb := Query()
		assert.Equal(t, "", qb.Build())
	})

	t.Run("All", func(t *testing.T) {
		t.Parallel()
		qb := Query().
			Equals("field1", "value1").
			NotEquals("field2", "value2").
			GreaterThan("field3", "10").
			LessThan("field4", "5").
			In("field5", []string{"val1", "val2"}).
			Exists("field6", true)
		assert.Equal(t, "where%5Bfield1%5D%5Bequals%5D=value1&where%5Bfield2%5D%5Bnot_equals%5D=value2&where%5Bfield3%5D%5Bgreater_than%5D=10&where%5Bfield4%5D%5Bless_than%5D=5&where%5Bfield5%5D%5Bin%5D=val1%2Cval2&where%5Bfield6%5D%5Bexists%5D=true", qb.Build())
	})
}
