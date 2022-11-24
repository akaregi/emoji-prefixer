package affix

type Property struct {
	Prefix string
	Suffix string
}

func NewAffixProps(prefix *string, suffix *string) Property {
	if *prefix == "" {
		*prefix = "regional_indicator_"
	}

	return Property{
		Prefix: *prefix,
		Suffix: *suffix,
	}
}
