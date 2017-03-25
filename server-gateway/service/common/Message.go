package common

type User string
type Number int

type Submission struct {
	User   User
	Number Number
}

func NewSubmission(user string, n int) *Submission {
	return &Submission{
		User:   User(user),
		Number: Number(n),
	}
}

func (s *Submission) Update(n Number) {
	s.Number += n
}

func (s *Submission) GetNumber() int {
	return int(s.Number)
}
