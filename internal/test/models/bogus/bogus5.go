package bogus

//go:generate reform

// reform:bogus
type Bogus5 struct {
	Bogus string `reform:"bogus,foo"` // field with "reform:" tag with unexpected value should generate error
}
