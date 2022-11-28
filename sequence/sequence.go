package sequence

type Sequence struct {
	Occurrence int32
	Word       string
}

type SequenceList []Sequence

func (p SequenceList) Len() int {
	return len(p)
}

func (p SequenceList) Less(i, j int) bool {
	return p[i].Occurrence < p[j].Occurrence
}

func (p SequenceList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
