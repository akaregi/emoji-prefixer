package printer

type Property struct {
	Verbose   bool
	Delimiter string
}

func NewPrinterProps(verbose *bool, delimiter *string) Property {
	return Property{
		Verbose:   *verbose,
		Delimiter: *delimiter,
	}
}
