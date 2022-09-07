package validations

type CommentValida struct {
	Name    string `valid:"Required;"`
	Email   string `valid:"Required;Email;"`
	Content string `valid:"Required;"`
	Aid     int    `valid:"Required;"`
	Pid     int    `valid:"Required;"`
}
