package util

type ArrayFlags []string

func (f *ArrayFlags) String() string {
	return ""
}

func (f *ArrayFlags) Set(v string) error {
	*f = append(*f, v)
	return nil
}
