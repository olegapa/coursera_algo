package rope_string

type Rope struct {
	left, right *Rope
	weight      int
	value       string
}

func NewRope(s string) *Rope {
	return &Rope{value: s, weight: len(s)}
}

func Concat(left, right *Rope) *Rope {
	return &Rope{
		left:   left,
		right:  right,
		weight: left.Length(),
	}
}

func (r *Rope) Length() int {
	if r == nil {
		return 0
	}
	if r.left == nil && r.right == nil {
		return len(r.value)
	}
	return r.weight + r.right.Length()
}

func (r *Rope) Index(i int) byte {
	if r.left == nil && r.right == nil {
		return r.value[i]
	}
	if i < r.weight {
		return r.left.Index(i)
	}
	return r.right.Index(i - r.weight)
}

func (r *Rope) ToString() string {
	if r.left == nil && r.right == nil {
		return r.value
	}
	return r.left.ToString() + r.right.ToString()
}
