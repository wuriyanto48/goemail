package goemail

// EmailSender interface
type EmailSender interface {
	Send() error
	SetTemplate(templateFile string, data interface{}) error
}
