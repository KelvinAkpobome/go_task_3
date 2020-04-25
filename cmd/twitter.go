package twitter


type Twitter struct {
	URL       string
	Username  string
	Followers int
}


func (t *Twitter) Feed() []string {
	return []string{
		"Twitter feeds",
		"Coding is basically copying other people's work",
	}
}

func (t *Twitter) Fame() int {
	return t.Followers
}
